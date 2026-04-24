package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
	"github.com/shopspring/decimal"
)

type InvoiceService struct {
	invoiceRepo     ports.InvoiceRepository
	installmentRepo ports.FeeInstallmentRepository
}

func NewInvoiceService(invRepo ports.InvoiceRepository, instRepo ports.FeeInstallmentRepository) *InvoiceService {
	return &InvoiceService{
		invoiceRepo:     invRepo,
		installmentRepo: instRepo,
	}
}

type CreateInvoiceCommand struct {
	TenantID         uuid.UUID
	StudentID        uuid.UUID
	FeeInstallmentID uuid.UUID
	Amount           string // peut être vide => on prend le montant de l'échéance
	Currency         financial.Currency
	DueDate          string // YYYY-MM-DD
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, cmd CreateInvoiceCommand) (*financial.Invoice, error) {
	// Récupérer l'échéance pour avoir les infos par défaut
	installment, err := s.installmentRepo.GetByID(ctx, cmd.FeeInstallmentID)
	if err != nil || installment == nil {
		return nil, fmt.Errorf("installment not found")
	}

	var amount decimal.Decimal
	if cmd.Amount != "" {
		amount, err = decimal.NewFromString(cmd.Amount)
		if err != nil {
			return nil, fmt.Errorf("invalid amount")
		}
	} else {
		amount = installment.Amount
	}

	currency := installment.Currency
	if cmd.Currency != "" {
		currency = cmd.Currency
	}

	dueDate := installment.DueDate
	if cmd.DueDate != "" {
		dueDate, _ = time.Parse("2006-01-02", cmd.DueDate)
	}

	invoice := &financial.Invoice{
		ID:               uuid.New(),
		TenantID:         cmd.TenantID,
		StudentID:        cmd.StudentID,
		FeeInstallmentID: cmd.FeeInstallmentID,
		InvoiceNumber:    generateInvoiceNumber(), // fonction utilitaire
		TotalAmount:      amount,
		Currency:         currency,
		Status:           financial.InvoiceStatusDraft,
		IssuedDate:       time.Now(),
		DueDate:          dueDate,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := s.invoiceRepo.Create(ctx, invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}

func generateInvoiceNumber() string {
	// simple : INV-YYYYMMDD-xxxx
	now := time.Now()
	return fmt.Sprintf("INV-%s-%d", now.Format("20060102"), now.UnixNano()%10000)
}
