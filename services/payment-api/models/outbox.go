package models

import (
	"time"

	"github.com/google/uuid"
)

type OutboxEvent struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Topic     string    `gorm:"type:varchar(100);not null"`
	Payload   []byte    `gorm:"type:jsonb;not null"`
	Processed bool      `gorm:"default:false;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
//da
