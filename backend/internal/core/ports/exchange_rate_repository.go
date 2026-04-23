package ports

import (
    "context"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
)

type ExchangeRateRepository interface {
    GetLatestRate(ctx context.Context, from, to financial.Currency) (decimal.Decimal, error)
    Create(ctx context.Context, rate *financial.ExchangeRate) error
}