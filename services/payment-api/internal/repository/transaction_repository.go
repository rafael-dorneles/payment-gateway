package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rafael-dorneles/payment-gateway/services/payment-api/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreatePaymentWithOutbox(ctx context.Context, tx *models.Transaction, event *models.OutboxEvent) error
	UpdateStatus(ctx context.Context, id uuid.UUID, oldStatus, newStatus string) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error)
	FindByUserId(ctx context.Context, userID uuid.UUID) ([]models.Transaction, error)
	FindByIDIdempotencyKey(ctx context.Context, key string) (*models.Transaction, error)
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

func (r *pgTransactionRepository) UpdateStatus(ctx context.Context, id uuid.UUID, oldStatus string, newStatus string) error {

	result := r.db.WithContext(ctx).Model(&models.Transaction{}).Where("id = ? AND status = ?", id, oldStatus).Update("status", newStatus)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("transition not allowed: current status is different from expected")
	}

	return nil

}

func (r *pgTransactionRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {

	var tx models.Transaction

	err := r.db.WithContext(ctx).First(&tx, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &tx, nil

}

func (r *pgTransactionRepository) FindByIDIdempotencyKey(ctx context.Context, idempotencyKey string) (*models.Transaction, error) {

	var tx models.Transaction

	err := r.db.WithContext(ctx).First(&tx, "idempotency_key = ?", idempotencyKey).Error

	if err != nil {
		return nil, err
	}

	return &tx, nil

}

func (r *pgTransactionRepository) FindByUserId(ctx context.Context, userID uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
