package models

import (
	"time"

	"github.com/google/uuid"
)

// cofre
type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index;unique"`
	Balance   int64     `gorm:"not null"` // Saldo em centavos
	Currency  string    `gorm:"not null"`
	UpdatedAt time.Time
}

// diario contabil
type LedgerEntry struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	FromAccountID uuid.UUID `gorm:"type:uuid;index"`
	ToAccountID   uuid.UUID `gorm:"type:uuid;index"`
	Amount        int64     `gorm:"not null"`
	TransactionID uuid.UUID `gorm:"type:uuid;index"`
	CreatedAt     time.Time
}

// Papel: Em sistemas financeiros, você nunca apenas
// altera o saldo. Você precisa registrar de onde saiu e
// para onde foi. Se o saldo do usuário está errado,
// você reconstrói o histórico olhando para as Entries.
