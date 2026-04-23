package services

import (
    "context"
    "fmt"

    "github.com/google/uuid"
)

// InvoiceGenerationService sera développé plus tard.
// Pour l'instant, il fournit juste une structure vide pour satisfaire les imports dans app.go.
type InvoiceGenerationService struct{}

func NewInvoiceGenerationService() *InvoiceGenerationService {
    return &InvoiceGenerationService{}
}

func (s *InvoiceGenerationService) GenerateInvoicesForAcademicYear(ctx context.Context, tenantID, academicYearID uuid.UUID) error {
    fmt.Println("Invoice generation not yet implemented")
    return nil
}