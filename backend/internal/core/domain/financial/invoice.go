package financial

import (
    "time"

    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

type Currency string

const (
    CDF Currency = "CDF"
    USD Currency = "USD"
)

type InvoiceStatus string

const (
    InvoiceStatusDraft         InvoiceStatus = "draft"
    InvoiceStatusSent          InvoiceStatus = "sent"
    InvoiceStatusPartiallyPaid InvoiceStatus = "partially_paid"
    InvoiceStatusPaid          InvoiceStatus = "paid"
    InvoiceStatusOverdue       InvoiceStatus = "overdue"
)

type Invoice struct {
    ID               uuid.UUID
    TenantID         uuid.UUID
    StudentID        uuid.UUID
    FeeInstallmentID uuid.UUID
    InvoiceNumber    string
    TotalAmount      decimal.Decimal
    Currency         Currency
    Status           InvoiceStatus
    IssuedDate       time.Time
    DueDate          time.Time
    CreatedAt        time.Time
    UpdatedAt        time.Time
}