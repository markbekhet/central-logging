package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/markbekhet/central-logging/logging-server/log"
	"google.golang.org/grpc"
)

const port = ":5051"

type server struct {
	pb.UnimplementedLoggerServer
}

func (s *server) Send(ctx context.Context, in *pb.HTTPLog) (*pb.LogRes, error) {
	// TODO: insert in the db
	fmt.Println("Received GRPC call")
	return &pb.LogRes{Status: 200}, nil
}

func main() {
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
