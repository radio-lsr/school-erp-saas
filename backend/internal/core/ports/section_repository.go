package ports

import (
    "context"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/academic"
)

type SectionRepository interface {
    Create(ctx context.Context, section *academic.Section) error
    GetByID(ctx context.Context, id uuid.UUID) (*academic.Section, error)
    ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*academic.Section, error)
    Update(ctx context.Context, section *academic.Section) error
    Delete(ctx context.Context, id uuid.UUID) error
}