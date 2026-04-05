package payment

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafael-dorneles/payment-gateway/internal/models"
)

type MockRepository struct {
	OnCreate func(t *models.Transaction) error
}

func (m *MockRepository) GetById(id uuid.UUID) (*models.Transaction, error) {
	panic("unimplemented")
}

func (m *MockRepository) UpdateStatus(id uuid.UUID, novoStatus string) error {
	panic("unimplemented")
}

func (m *MockRepository) Create(t *models.Transaction) error {
	return m.OnCreate(t)
}

func TestPayemntService_create(t *testing.T) {

	t.Run("Deve retornar erro se o valor for zero", func(t *testing.T) {
		mockRepo := &MockRepository{}
		service := NewPaymentService(mockRepo)
		req := models.PaymentRequest{Amount: 0}

		_, err := service.Create(req)

		if err == nil {
			t.Error("Esperava um erro para valor zero, mas recebi nil")
		}
	})

	t.Run("Deve criar transação com sucesso quando o valor for positivo", func(t *testing.T) {
		mockRepo := &MockRepository{
			OnCreate: func(t *models.Transaction) error {
				return nil // Simula sucesso no banco
			},
		}
		service := NewPaymentService(mockRepo)
		req := models.PaymentRequest{Amount: 100}

		tx, err := service.Create(req)

		if err != nil {
			t.Errorf("Não esperava erro para valor positivo, mas recebi: %v", err)
		}

		if tx == nil {
			t.Fatal("Transação retornada é nula")
		}

		if tx.Status != "pending" {
			t.Errorf("Status esperado: pending, recebido: %s", tx.Status)
		}
	})
}
