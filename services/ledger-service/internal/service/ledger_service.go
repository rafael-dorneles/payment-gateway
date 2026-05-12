package service

import (
	"context"

	"payment-gateway/services/ledger-service/internal/dto"
	"payment-gateway/services/ledger-service/internal/repository"
)

type LedgerService interface {
	ProcessPayment(ctx context.Context, event dto.PaymentEvent) error
}

type ledgerService struct {
	repo repository.TransactionRepository
}

func NewLedgerService(repo repository.TransactionRepository) LedgerService {
	return &ledgerService{
		repo: repo,
	}
}

func (s *ledgerService) ProcessPayment(ctx context.Context, event dto.PaymentEvent)
