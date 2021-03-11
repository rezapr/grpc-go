package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rezapr/grpc-go/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello")

	conn, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error starting TCP/50051 :", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(conn); err != nil {
		log.Fatalf("Error to serve: %v", err)
	}
}
