package ports

import (
    "context"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
)

type PaymentRepository interface {
    Create(ctx context.Context, p *financial.Payment) error
    GetByID(ctx context.Context, id uuid.UUID) (*financial.Payment, error)
    GetByTransactionID(ctx context.Context, transactionID string) (*financial.Payment, error)
    GetTotalPaidForInvoice(ctx context.Context, invoiceID uuid.UUID, currency financial.Currency) (decimal.Decimal, error)
}