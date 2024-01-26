package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/markbekhet/central-logging/log"
	"google.golang.org/grpc"
)

func main() {
	grpcAddr := os.Getenv("grpcAddr")
	grpcPort := os.Getenv("grpcPort")
	addr := grpcAddr + ":" + grpcPort
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(-1)
	}
	defer conn.Close()
	c := pb.NewLoggerClient(conn)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time := time.Now()
		fmt.Println(r.RemoteAddr)
		if r.Method == http.MethodGet {
			w.Write([]byte("Hello"))
			hostname, _ := os.Hostname()
			c.Send(context.Background(), &pb.HTTPLog{
				StatusCode: http.StatusOK,
				ServerHost: hostname,
				UserIP:     r.RemoteAddr,
				Date: &pb.HTTPLog_Date{
					Day:    int32(time.Day()),
					Month:  int32(time.Month()),
					Year:   int32(time.Year()),
					Hour:   int32(time.Hour()),
					Minute: int32(time.Minute()),
					Second: int32(time.Second()),
					Nano:   int64(time.Nanosecond()),
				},
			})
		}
	})
	http.ListenAndServe(":8080", nil)

}
