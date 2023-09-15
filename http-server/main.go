package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/markbekhet/central-logging/http-server/log"
	"google.golang.org/grpc"
)

const grpcAddr = "localhost:5051"

func main() {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLoggerClient(conn)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time := time.Now()
		fmt.Println(r.RemoteAddr)
		if r.Method == http.MethodGet {
			w.Write([]byte("Hello"))
			c.Send(context.Background(), &pb.HTTPLog{
				StatusCode: http.StatusOK,
				ServerIP:   r.Host,
				UserIP:     r.RemoteAddr,
				Date: &pb.HTTPLog_Date{
					Day:   int32(time.Day()),
					Month: int32(time.Month()),
					Year:  int32(time.Year()),
				},
			})
		}
	})
	http.ListenAndServe(":8080", nil)

}
