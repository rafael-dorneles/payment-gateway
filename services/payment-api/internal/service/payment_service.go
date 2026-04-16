package service

import (
	"context"
	"encoding/json"
	"errors"
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
