package main

import (
	"log"
	"net"

	"github.com/mohamedfawas/calculator-service/endpoint"
	pb "github.com/mohamedfawas/calculator-service/proto"
	"github.com/mohamedfawas/calculator-service/service"
	"github.com/mohamedfawas/calculator-service/transport"
	"google.golang.org/grpc"
	// Optional: helps with debugging
)

func main() {
	// Create a listener on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create service and endpoints
	svc := service.NewCalculatorService()
	endpoints := endpoint.MakeEndpoints(svc)

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, transport.NewGRPCServer(endpoints))

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
