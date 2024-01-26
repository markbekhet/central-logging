package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/markbekhet/central-logging/log"
	"google.golang.org/grpc"
)

const port = ":5051"

type server struct {
	pb.UnimplementedLoggerServer
}

func (s *server) Send(ctx context.Context, in *pb.HTTPLog) (*pb.LogRes, error) {
	// TODO: insert in the db
	err := AddHTTPLOG(in)
	if err != nil {
		log.Fatalln(err)
		return &pb.LogRes{Status: 503}, nil
	}
	return &pb.LogRes{Status: 200}, nil
}

func main() {
	fmt.Println(os.Environ())
	ConnectDB(os.Getenv("dbHost"), os.Getenv("dbUser"),
		os.Getenv("dbPassword"), "postgres", 5432)
	err := CreateHTTPLogTable()
	if err != nil {
		os.Exit(-1)
	}
	defer CloseDB()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
