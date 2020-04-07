package main

import (
	"google.golang.org/grpc"
	pb "health/service/clinical"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMeasureServiceServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
