package ports

import (
    "context"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/student"
)

type StudentRepository interface {
    Create(ctx context.Context, s *student.Student) error
    GetByID(ctx context.Context, id uuid.UUID) (*student.Student, error)
    ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*student.Student, error)
    Update(ctx context.Context, s *student.Student) error
    Delete(ctx context.Context, id uuid.UUID) error
}