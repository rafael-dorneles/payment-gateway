package repository

import (
	"context"

	"github.com/rafael-dorneles/payment-gateway/services/payment-api/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreatePaymentWithOutbox(ctx context.Context, tx *models.Transaction, event *models.OutboxEvent) error
}

type pgTransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &pgTransactionRepository{db: db}
}

func (r *pgTransactionRepository) CreatePaymentWithOutbox(ctx context.Context, tx *models.Transaction, event *models.OutboxEvent) error {
	return r.db.WithContext(ctx).Transaction(func(dbTx *gorm.DB) error {

		if err := dbTx.Create(tx).Error; err != nil {
			return err
		}

		if err := dbTx.Create(event).Error; err != nil {
			return err
		}

		return nil
	})
}
