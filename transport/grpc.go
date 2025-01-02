package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"                                    // Add this import
	localendpoint "github.com/mohamedfawas/calculator-service/endpoint" // Rename to avoid conflict
	pb "github.com/mohamedfawas/calculator-service/proto"
)

type grpcServer struct {
	add      endpoint.Endpoint
	subtract endpoint.Endpoint
	pb.UnimplementedCalculatorServer
}

func NewGRPCServer(endpoints localendpoint.Endpoints) pb.CalculatorServer {
	return &grpcServer{
		add:      endpoints.Add,
		subtract: endpoints.Subtract,
	}
}

func (s *grpcServer) Add(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	resp, err := s.add(ctx, localendpoint.CalculateRequest{A: req.A, B: req.B})
	if err != nil {
		return nil, err
	}
	return &pb.CalculateResponse{Result: resp.(localendpoint.CalculateResponse).Result}, nil
}

func (s *grpcServer) Subtract(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	resp, err := s.subtract(ctx, localendpoint.CalculateRequest{A: req.A, B: req.B})
	if err != nil {
		return nil, err
	}
	return &pb.CalculateResponse{Result: resp.(localendpoint.CalculateResponse).Result}, nil
}
