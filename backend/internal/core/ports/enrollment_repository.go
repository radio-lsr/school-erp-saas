package ports

import (
    "context"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/enrollment"
)

type EnrollmentRepository interface {
    Create(ctx context.Context, e *enrollment.Enrollment) error
    GetByID(ctx context.Context, id uuid.UUID) (*enrollment.Enrollment, error)
    ListByStudent(ctx context.Context, studentID uuid.UUID) ([]*enrollment.Enrollment, error)
    ListBySection(ctx context.Context, sectionID uuid.UUID) ([]*enrollment.Enrollment, error)
    CountActiveBySection(ctx context.Context, sectionID uuid.UUID) (int, error)
    Update(ctx context.Context, e *enrollment.Enrollment) error
    Delete(ctx context.Context, id uuid.UUID) error
    ListActiveByAcademicYear(ctx context.Context, tenantID, academicYearID uuid.UUID) ([]*enrollment.Enrollment, error)
}