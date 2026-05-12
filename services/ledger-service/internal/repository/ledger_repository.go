package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rafael-dorneles/payment-gateway/services/ledger-service/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LedgerRepository struct {
	db *gorm.DB
}

// func NewLedgerRepository(db *gorm.DB) TransactionRepository {
// 	return &pgLedgerRepository{db: db}
// }

func (r *LedgerRepository) Transfer(ctx context.Context, fromID, toID uuid.UUID, amount int64, txID uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var fromAccount models.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&fromAccount, "id = ?", fromID).Error; err != nil {
			return err
		}

		if fromAccount.Balance < amount {
			return errors.New("insufficient funds")
		}
		if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance-amount).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Account{}).Where("id = ?", toID).UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		entry := models.LedgerEntry{
			ID:            uuid.New(),
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
			TransactionID: txID,
		}

		return tx.Create(&entry).Error
	})
}
