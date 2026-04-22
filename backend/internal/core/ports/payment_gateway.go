package ports

import (
	"context"

	"github.com/shopspring/decimal"
)

type PaymentGateway interface {
	InitiatePayment(ctx context.Context, req PaymentRequest) (*PaymentResponse, error)
	CheckPaymentStatus(ctx context.Context, transactionID string) (*PaymentStatus, error)
}

type PaymentRequest struct {
	Amount      decimal.Decimal
	Currency    string
	PhoneNumber string
	Description string
	Reference   string
	CallbackURL string
}

type PaymentResponse struct {
	TransactionID string
	Status        string
	RedirectURL   string // pour les paiements web
}
