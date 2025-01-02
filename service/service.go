package service

import "context"

// Service defines the calculator service interface
type CalculatorService interface {
	Add(ctx context.Context, a, b int32) (int32, error)
	Subtract(ctx context.Context, a, b int32) (int32, error)
}

type calculatorService struct{}

// NewCalculatorService creates a new calculator service
func NewCalculatorService() CalculatorService {
	return &calculatorService{}
}

// Add performs addition
func (s *calculatorService) Add(ctx context.Context, a, b int32) (int32, error) {
	return a + b, nil
}

// Subtract performs subtraction
func (s *calculatorService) Subtract(ctx context.Context, a, b int32) (int32, error) {
	return a - b, nil
}
