package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
	"github.com/shopspring/decimal"
)

type PostgresExchangeRateRepository struct {
	db *pgxpool.Pool
}

func NewPostgresExchangeRateRepository(db *pgxpool.Pool) ports.ExchangeRateRepository {
	return &PostgresExchangeRateRepository{db: db}
}

func (r *PostgresExchangeRateRepository) GetLatestRate(ctx context.Context, from, to financial.Currency) (decimal.Decimal, error) {
	query := `SELECT rate FROM exchange_rates WHERE from_currency = $1 AND to_currency = $2 ORDER BY effective_date DESC LIMIT 1`
	var rate decimal.Decimal
	err := r.db.QueryRow(ctx, query, string(from), string(to)).Scan(&rate)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return decimal.NewFromInt(1), nil // fallback 1:1
		}
		return decimal.Zero, err
	}
	return rate, nil
}

func (r *PostgresExchangeRateRepository) Create(ctx context.Context, rate *financial.ExchangeRate) error {
	query := `INSERT INTO exchange_rates (id, tenant_id, from_currency, to_currency, rate, effective_date)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, rate.ID, rate.TenantID, rate.FromCurrency, rate.ToCurrency, rate.Rate, rate.EffectiveDate)
	return err
}
