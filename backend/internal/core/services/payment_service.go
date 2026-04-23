package services

import (
    "context"
    "errors"
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PaymentService struct {
    invoiceRepo    ports.InvoiceRepository
    paymentRepo    ports.PaymentRepository
    exchangeRepo   ports.ExchangeRateRepository
}

func NewPaymentService(invoiceRepo ports.InvoiceRepository, paymentRepo ports.PaymentRepository, exchangeRepo ports.ExchangeRateRepository) *PaymentService {
    return &PaymentService{
        invoiceRepo:  invoiceRepo,
        paymentRepo:  paymentRepo,
        exchangeRepo: exchangeRepo,
    }
}

type AddPaymentCommand struct {
    InvoiceID     uuid.UUID
    AmountPaid    decimal.Decimal
    CurrencyPaid  financial.Currency
    PaymentMethod string
    Reference     string
}

func (s *PaymentService) AddPayment(ctx context.Context, cmd AddPaymentCommand) (*financial.Payment, error) {
    invoice, err := s.invoiceRepo.GetByID(ctx, cmd.InvoiceID)
    if err != nil || invoice == nil {
        return nil, errors.New("invoice not found")
    }

    totalPaid, err := s.paymentRepo.GetTotalPaidForInvoice(ctx, invoice.ID, invoice.Currency)
    if err != nil {
        return nil, err
    }

    paidAmountInInvoiceCurrency := cmd.AmountPaid
    exchangeRate := decimal.NullDecimal{}

    if cmd.CurrencyPaid != invoice.Currency {
        rate, err := s.exchangeRepo.GetLatestRate(ctx, cmd.CurrencyPaid, invoice.Currency)
        if err != nil {
            return nil, errors.New("exchange rate not available")
        }
        paidAmountInInvoiceCurrency = cmd.AmountPaid.Mul(rate)
        exchangeRate = decimal.NewNullDecimal(rate)
    }

    remaining := invoice.TotalAmount.Sub(totalPaid)
    if paidAmountInInvoiceCurrency.GreaterThan(remaining) {
        return nil, errors.New("payment exceeds remaining balance")
    }

    payment := &financial.Payment{
        ID:            uuid.New(),
        TenantID:      invoice.TenantID,
        InvoiceID:     cmd.InvoiceID,
        AmountPaid:    cmd.AmountPaid,
        CurrencyPaid:  cmd.CurrencyPaid,
        PaymentDate:   time.Now(),
        PaymentMethod: cmd.PaymentMethod,
        Reference:     cmd.Reference,
        ExchangeRate:  exchangeRate,
        CreatedAt:     time.Now(),
    }

    if err := s.paymentRepo.Create(ctx, payment); err != nil {
        return nil, err
    }

    newTotalPaid := totalPaid.Add(paidAmountInInvoiceCurrency)
    if newTotalPaid.GreaterThanOrEqual(invoice.TotalAmount) {
        invoice.Status = financial.InvoiceStatusPaid
    } else {
        invoice.Status = financial.InvoiceStatusPartiallyPaid
    }
    _ = s.invoiceRepo.Update(ctx, invoice)

    return payment, nil
}