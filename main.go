package main

import (
	"fmt"
	"log"
	"net"

	"github.com/heathesh/grpc-helloworld/healthcheck"
	"google.golang.org/grpc"
)

func main() {
	portNumber := ":7500"
	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register your gRPC server implementation
	healthcheck.RegisterHealthCheckServer(grpcServer, &healthcheck.Server{})

	// Start the gRPC server
	fmt.Println("Starting gRPC server on port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
