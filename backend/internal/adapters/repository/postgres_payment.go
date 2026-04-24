package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
	"github.com/shopspring/decimal"
)

type PostgresPaymentRepository struct {
	db *pgxpool.Pool
}

func NewPostgresPaymentRepository(db *pgxpool.Pool) ports.PaymentRepository {
	return &PostgresPaymentRepository{db: db}
}

func (r *PostgresPaymentRepository) Create(ctx context.Context, p *financial.Payment) error {
	query := `INSERT INTO payments (id, tenant_id, invoice_id, amount_paid, currency_paid, payment_date, payment_method, reference, exchange_rate, created_at)
              VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	_, err := r.db.Exec(ctx, query,
		p.ID, p.TenantID, p.InvoiceID, p.AmountPaid, p.CurrencyPaid, p.PaymentDate,
		p.PaymentMethod, p.Reference, p.ExchangeRate, p.CreatedAt)
	return err
}

func (r *PostgresPaymentRepository) GetByID(ctx context.Context, id uuid.UUID) (*financial.Payment, error) {
	query := `SELECT id, tenant_id, invoice_id, amount_paid, currency_paid, payment_date, payment_method, reference, exchange_rate, created_at
              FROM payments WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	var p financial.Payment
	err := row.Scan(&p.ID, &p.TenantID, &p.InvoiceID, &p.AmountPaid, &p.CurrencyPaid,
		&p.PaymentDate, &p.PaymentMethod, &p.Reference, &p.ExchangeRate, &p.CreatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *PostgresPaymentRepository) GetByTransactionID(ctx context.Context, transactionID string) (*financial.Payment, error) {
	query := `SELECT id, tenant_id, invoice_id, amount_paid, currency_paid, payment_date, payment_method, reference, exchange_rate, created_at
              FROM payments WHERE reference = $1` // en supposant que la reference est le transactionID
	row := r.db.QueryRow(ctx, query, transactionID)
	var p financial.Payment
	err := row.Scan(&p.ID, &p.TenantID, &p.InvoiceID, &p.AmountPaid, &p.CurrencyPaid,
		&p.PaymentDate, &p.PaymentMethod, &p.Reference, &p.ExchangeRate, &p.CreatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *PostgresPaymentRepository) GetTotalPaidForInvoice(ctx context.Context, invoiceID uuid.UUID, currency financial.Currency) (decimal.Decimal, error) {
	query := `SELECT COALESCE(SUM(amount_paid), 0) FROM payments WHERE invoice_id = $1 AND currency_paid = $2`
	var total decimal.Decimal
	err := r.db.QueryRow(ctx, query, invoiceID, string(currency)).Scan(&total)
	return total, err
}
