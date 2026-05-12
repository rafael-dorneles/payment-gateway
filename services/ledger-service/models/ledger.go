package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index;unique"`
	Balance   int64     `gorm:"not null"` // Saldo em centavos
	Currency  string    `gorm:"not null"`
	UpdatedAt time.Time
}

type LedgerEntry struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	FromAccountID uuid.UUID `gorm:"type:uuid;index"`
	ToAccountID   uuid.UUID `gorm:"type:uuid;index"`
	Amount        int64     `gorm:"not null"`
	TransactionID uuid.UUID `gorm:"type:uuid;index"`
	CreatedAt     time.Time
}
