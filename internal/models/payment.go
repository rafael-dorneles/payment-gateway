package models

import (
	"time"
)

// pagamento no sistema
type Transaction struct {
	ID        string    `json:"id"`
	Amout     float64   `json:"amout"`
	Currency  float64   `json:"currency"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_id"`
	UpdatedAt time.Time `json:"updated_id"`
}

// oq a api vai receber
type PaytmentRequest struct {
	CardNumber     string  `json:"card_number"`
	CardName       string  `json:"card_name"`
	ExpirationDate string  `json:"expiration_date"`
	CVV            string  `json:"cvv"`
	Amount         float64 `json:"amount"`
}
