package main

import (
	"log"
	"net"

	pb "github.com/markbekhet/central-logging/log"
	"google.golang.org/grpc"
)

const port = ":5051"

type server struct{}

func main() {
	_, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	log.Println("Server listening on port %s", port)

}
