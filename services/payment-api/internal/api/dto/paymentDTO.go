package dto

import "time"

type PaymentRequest struct {
	UserID   string `json: "user_id" validate:"required,uuid`
	Amount   int64  `json: "amount" validate : "required, gt-0"`
	Currency string `json: "currency" validate : "required, len=3"`
}

type PaymentResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	Amount        int64     `json:"amount"`     // Adicionado
	Currency      string    `json:"currency"`   // Adicionado
	CreatedAt     time.Time `json:"created_at"` // Adicionado
	Message       string    `json:"message"`
}
