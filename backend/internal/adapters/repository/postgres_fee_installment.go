package repository

import (
    "context"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresFeeInstallmentRepository struct {
    db *pgxpool.Pool
}

func NewPostgresFeeInstallmentRepository(db *pgxpool.Pool) ports.FeeInstallmentRepository {
    return &PostgresFeeInstallmentRepository{db: db}
}

func (r *PostgresFeeInstallmentRepository) Create(ctx context.Context, i *financial.FeeInstallment) error { return nil }
func (r *PostgresFeeInstallmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*financial.FeeInstallment, error) { return nil, nil }
func (r *PostgresFeeInstallmentRepository) ListByFeeStructure(ctx context.Context, feeStructureID uuid.UUID) ([]*financial.FeeInstallment, error) { return nil, nil }
func (r *PostgresFeeInstallmentRepository) Update(ctx context.Context, i *financial.FeeInstallment) error { return nil }
func (r *PostgresFeeInstallmentRepository) Delete(ctx context.Context, id uuid.UUID) error { return nil }