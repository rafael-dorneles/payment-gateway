package payment

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rafael-dorneles/payment-gateway/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(t *models.Transaction) error
	GetById(id uuid.UUID) (*models.Transaction, error)
	UpdateStatus(id uuid.UUID, novoStatus string) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) Create(t *models.Transaction) error {
	return r.db.Create(t).Error
}

func (r *paymentRepository) UpdateStatus(id uuid.UUID, novoStatus string) error {
	result := r.db.Model(&models.Transaction{}).
		Where("id = ?", id).
		Update("status", novoStatus)

	if result.Error != nil {
		return result.Error
	}

	// Verificamos se o ID realmente existia no banco
	if result.RowsAffected == 0 {
		return errors.New("pagamento não encontrado")
	}

	return nil
}

//o r* paymentRepository significa que esta funcao pertence a sua struct do repositorio

// recebe um uiid.UUID e um identificador unico universal

//devolve um ponteiro

func (r *paymentRepository) GetById(id uuid.UUID) (*models.Transaction, error) {
	var tx models.Transaction

	err := r.db.First(&tx, "id =: ?", id).Error

	if err != nil {
		return nil, err
	}

	return &tx, nil
}
