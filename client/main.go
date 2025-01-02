package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mohamedfawas/calculator-service/proto"
	"google.golang.org/grpc"
	// For insecure connection
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test addition
	addResp, err := client.Add(ctx, &pb.CalculateRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Addition: 10 + 5 = %d", addResp.Result)

	// Test subtraction
	subResp, err := client.Subtract(ctx, &pb.CalculateRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Subtraction: 10 - 5 = %d", subResp.Result)
}
