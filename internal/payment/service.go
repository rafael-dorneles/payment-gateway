package payment

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rafael-dorneles/payment-gateway/internal/models"
)

type PaymentService interface {
	Create(record models.PaymentRequest) (*models.Transaction, error)
	GetById(id uuid.UUID) (*models.Transaction, error)
}

type paymentService struct {
	repo PaymentRepository
}

func NewPaymentService(r PaymentRepository) PaymentService {
	return &paymentService{
		repo: r,
	}
}

func (s *paymentService) Create(record models.PaymentRequest) (*models.Transaction, error) {
	if record.Amount <= 0 {
		return nil, errors.New("o valor deve ser maior que zero")
	}

	tx := &models.Transaction{
		ID:       uuid.New(),
		Amount:   record.Amount,
		Status:   "pending",
		Currency: "BRL",
	}

	if err := s.repo.Create(tx); err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *paymentService) GetById(id uuid.UUID) (*models.Transaction, error) {

	tx, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return tx, nil

}
