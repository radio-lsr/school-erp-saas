package repository

import (
    "context"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/shopspring/decimal"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresPaymentRepository struct {
    db *pgxpool.Pool
}

func NewPostgresPaymentRepository(db *pgxpool.Pool) ports.PaymentRepository {
    return &PostgresPaymentRepository{db: db}
}

func (r *PostgresPaymentRepository) Create(ctx context.Context, p *financial.Payment) error { return nil }
func (r *PostgresPaymentRepository) GetByID(ctx context.Context, id uuid.UUID) (*financial.Payment, error) { return nil, nil }
func (r *PostgresPaymentRepository) GetByTransactionID(ctx context.Context, transactionID string) (*financial.Payment, error) { return nil, nil }
func (r *PostgresPaymentRepository) GetTotalPaidForInvoice(ctx context.Context, invoiceID uuid.UUID, currency financial.Currency) (decimal.Decimal, error) { return decimal.Zero, nil }