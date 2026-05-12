package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/rafael-dorneles/payment-gateway/pkg/cache"
	"github.com/rafael-dorneles/payment-gateway/services/payment-api/internal/dto"
	"github.com/rafael-dorneles/payment-gateway/services/payment-api/internal/repository"
	"github.com/rafael-dorneles/payment-gateway/services/payment-api/models"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, input dto.PaymentRequest, idempotencyKey string) (*dto.PaymentResponse, error)
}

type paymentService struct {
	repo  repository.TransactionRepository
	cache *cache.Client
}

func NewPaymentService(repo repository.TransactionRepository, cache *cache.Client) PaymentService {
	return &paymentService{
		repo:  repo,
		cache: cache,
	}
}

func (s *paymentService) ProcessPayment(ctx context.Context, input dto.PaymentRequest, idempotencyKey string) (*dto.PaymentResponse, error) {
	cachedStatus, _ := s.cache.Get(ctx, idempotencyKey)
	if cachedStatus != "" {
		return nil, errors.New("request already processed or in progress")
	}

	txID := uuid.New()
	userID, _ := uuid.Parse(input.UserID)

	transaction := &models.Transaction{
		ID:             txID,
		UserID:         userID,
		Amount:         input.Amount,
		Currency:       input.Currency,
		Status:         "pending",
		IdempotencyKey: idempotencyKey,
	}

	payload, _ := json.Marshal(transaction)
	event := &models.OutboxEvent{
		ID:      uuid.New(),
		Topic:   "payments.created",
		Payload: payload,
	}

	err := s.repo.CreatePaymentWithOutbox(ctx, transaction, event)
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetIdempotencyKey(ctx, idempotencyKey, "processed", 24*time.Hour)

	return &dto.PaymentResponse{
		TransactionID: txID.String(),
		Status:        "pending",
		Message:       "Payment received and is being processed",
	}, nil
}

func (s *paymentService) UpdatePaymentStatus(ctx context.Context, id uuid.UUID, oldStatus string, newStatus string) error {

	tx, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if !s.isTransitionAllowed(tx.Status, newStatus) {
		return errors.New("transição de status inválida: de " + tx.Status + " para " + newStatus)
	}

	err = s.repo.UpdateStatus(ctx, id, tx.Status, newStatus)
	if err != nil {
		return err
	}

	return nil
}

func (s *paymentService) isTransitionAllowed(current, next string) bool {
	switch current {
	case "pending":
		return next == "approved" || next == "rejected" || next == "canceled"
	case "approved":
		return next == "refunded"
	case "rejected", "canceled", "refunded":
		return false
	default:
		return false
	}
}

func (s *paymentService) GetPaymentStatus(ctx context.Context, id uuid.UUID) (*dto.PaymentResponse, error) {

	tx, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &dto.PaymentResponse{
		TransactionID: tx.ID.String(),
		Status:        tx.Status,
		Message:       "Status atual da transacao",
	}, nil
}

func (s *paymentService) CancelPayment(ctx context.Context, id uuid.UUID, reason string) error {

	tx, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return err
	}

	log.Printf("[Service] Cancelando pagamento %s. Motivo: %s", id, reason)

	return s.UpdatePaymentStatus(ctx, id, tx.Status, "canceled")
}

func (s *paymentService) GetTransactionsByUserID(ctx context.Context, userID uuid.UUID) ([]dto.PaymentResponse, error) {
	transactions, err := s.repo.FindByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	var response []dto.PaymentResponse
	for _, tx := range transactions {
		response = append(response, dto.PaymentResponse{
			TransactionID: tx.ID.String(),
			Status:        tx.Status,
			Amount:        tx.Amount,
			Currency:      tx.Currency,
			CreatedAt:     tx.CreatedAt,
			Message:       "Transaction history record",
		})
	}

	if response == nil {
		return []dto.PaymentResponse{}, nil
	}

	return response, nil
}
