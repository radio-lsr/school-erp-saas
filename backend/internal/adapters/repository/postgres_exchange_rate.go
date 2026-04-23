package repository

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/shopspring/decimal"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresExchangeRateRepository struct {
    db *pgxpool.Pool
}

func NewPostgresExchangeRateRepository(db *pgxpool.Pool) ports.ExchangeRateRepository {
    return &PostgresExchangeRateRepository{db: db}
}

func (r *PostgresExchangeRateRepository) GetLatestRate(ctx context.Context, from, to financial.Currency) (decimal.Decimal, error) {
    // stub: retourne 1:1
    return decimal.NewFromInt(1), nil
}
func (r *PostgresExchangeRateRepository) Create(ctx context.Context, rate *financial.ExchangeRate) error { return nil }