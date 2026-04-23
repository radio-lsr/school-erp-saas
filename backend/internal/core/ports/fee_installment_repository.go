package ports

import (
    "context"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
)

type FeeInstallmentRepository interface {
    Create(ctx context.Context, i *financial.FeeInstallment) error
    GetByID(ctx context.Context, id uuid.UUID) (*financial.FeeInstallment, error)
    ListByFeeStructure(ctx context.Context, feeStructureID uuid.UUID) ([]*financial.FeeInstallment, error)
    Update(ctx context.Context, i *financial.FeeInstallment) error
    Delete(ctx context.Context, id uuid.UUID) error
}