package repository

import (
    "context"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresInvoiceRepository struct {
    db *pgxpool.Pool
}

func NewPostgresInvoiceRepository(db *pgxpool.Pool) ports.InvoiceRepository {
    return &PostgresInvoiceRepository{db: db}
}

func (r *PostgresInvoiceRepository) Create(ctx context.Context, invoice *financial.Invoice) error { return nil }
func (r *PostgresInvoiceRepository) GetByID(ctx context.Context, id uuid.UUID) (*financial.Invoice, error) { return nil, nil }
func (r *PostgresInvoiceRepository) GetByNumber(ctx context.Context, number string) (*financial.Invoice, error) { return nil, nil }
func (r *PostgresInvoiceRepository) ListByStudent(ctx context.Context, studentID uuid.UUID) ([]*financial.Invoice, error) { return nil, nil }
func (r *PostgresInvoiceRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*financial.Invoice, error) { return nil, nil }
func (r *PostgresInvoiceRepository) Update(ctx context.Context, invoice *financial.Invoice) error { return nil }