package ports

import (
    "context"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
)

type InvoiceRepository interface {
    Create(ctx context.Context, invoice *financial.Invoice) error
    GetByID(ctx context.Context, id uuid.UUID) (*financial.Invoice, error)
    GetByNumber(ctx context.Context, number string) (*financial.Invoice, error)
    ListByStudent(ctx context.Context, studentID uuid.UUID) ([]*financial.Invoice, error)
    ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*financial.Invoice, error)
    Update(ctx context.Context, invoice *financial.Invoice) error
}