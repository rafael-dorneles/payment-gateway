package domain

import "context"

type PaymentStrategy interface {
	Process(ctx context.Context, payment *Payment) (*PaymentResponse, error)
}

type PaymentResponse struct {
	TransactionID string
	Status        string
	GatewayMsg    string
	Payload       map[string]interface{}
}
