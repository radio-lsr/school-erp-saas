package services

import (
    "context"
    "fmt"
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type InvoiceGenerationService struct {
    enrollmentRepo ports.EnrollmentRepository
    feeStructRepo  ports.FeeStructureRepository
    installmentRepo ports.FeeInstallmentRepository
    invoiceRepo    ports.InvoiceRepository
}

func (s *InvoiceGenerationService) GenerateInvoicesForAcademicYear(ctx context.Context, tenantID, academicYearID uuid.UUID) error {
    // 1. Récupérer toutes les inscriptions actives pour l'année
    enrollments, err := s.enrollmentRepo.ListActiveByAcademicYear(ctx, tenantID, academicYearID)
    if err != nil {
        return err
    }

    for _, enrollment := range enrollments {
        // 2. Trouver la structure tarifaire pour le niveau de l'élève
        section, _ := ... // obtenir section pour avoir grade_level_id
        feeStruct, err := s.feeStructRepo.GetByGradeAndYear(ctx, tenantID, section.GradeLevelID, academicYearID)
        if err != nil {
            continue // peut-être pas défini
        }

        // 3. Récupérer les échéances
        installments, err := s.installmentRepo.ListByFeeStructure(ctx, feeStruct.ID)
        if err != nil {
            continue
        }

        // 4. Créer une facture par échéance
        for _, inst := range installments {
            invoice := &financial.Invoice{
                ID:               uuid.New(),
                TenantID:         tenantID,
                StudentID:        enrollment.StudentID,
                FeeInstallmentID: inst.ID,
                InvoiceNumber:    generateInvoiceNumber(),
                TotalAmount:      inst.Amount,
                Currency:         inst.Currency,
                Status:           financial.InvoiceStatusDraft,
                IssuedDate:       time.Now(),
                DueDate:          inst.DueDate,
            }
            if err := s.invoiceRepo.Create(ctx, invoice); err != nil {
                return err
            }
        }
    }
    return nil
}