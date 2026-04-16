package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID     uuid.UUID `gorm:"type: uuid ; primaryKey"`
	UserID uuid.UUID `gorm:"type: uuid ; index; not null"`

	Amount   int64  `gorm:"not null"`
	Currency string `gorm:"type:varchar(3); default:'BRL'"`

	Status string `gorm:"type:varchar(20);index"`

	IdempotencyKey string `gorm:"type : varchar(100); uniqueIndex; not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
