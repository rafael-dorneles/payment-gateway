package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ExternalID string    `gorm:"type:varchar(255);index" json:"external_id"`

	Amount int64 `gorm:"not null" json:"amount"`

	Currency string `gorm:"type:varchar(3);not null" json:"currency"`

	Status string `gorm:"type:varchar(20);not null" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// O que a API vai receber
type PaymentRequest struct {
	CardNumber     string `json:"card_number"`
	CardName       string `json:"card_name"`
	ExpirationDate string `json:"expiration_date"`
	CVV            string `json:"cvv"`
	Amount         int64  `json:"amount"`
	Currency       string `json:"currency"`
}
