package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mohamedfawas/calculator-service/service"
)

type Endpoints struct {
	Add      endpoint.Endpoint
	Subtract endpoint.Endpoint
}

type CalculateRequest struct {
	A int32
	B int32
}

type CalculateResponse struct {
	Result int32
}

func MakeEndpoints(s service.CalculatorService) Endpoints {
	return Endpoints{
		Add:      makeAddEndpoint(s),
		Subtract: makeSubtractEndpoint(s),
	}
}

func makeAddEndpoint(s service.CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculateRequest)
		result, err := s.Add(ctx, req.A, req.B)
		return CalculateResponse{Result: result}, err
	}
}

func makeSubtractEndpoint(s service.CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculateRequest)
		result, err := s.Subtract(ctx, req.A, req.B)
		return CalculateResponse{Result: result}, err
	}
}
