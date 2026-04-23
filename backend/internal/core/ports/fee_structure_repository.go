package ports

import (
    "context"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
)

type FeeStructureRepository interface {
    Create(ctx context.Context, f *financial.FeeStructure) error
    GetByID(ctx context.Context, id uuid.UUID) (*financial.FeeStructure, error)
    GetByGradeAndYear(ctx context.Context, tenantID, gradeLevelID, academicYearID uuid.UUID) (*financial.FeeStructure, error)
    ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*financial.FeeStructure, error)
    Update(ctx context.Context, f *financial.FeeStructure) error
    Delete(ctx context.Context, id uuid.UUID) error
}