package financial

import (
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

type Payment struct {
    ID            uuid.UUID
    TenantID      uuid.UUID
    InvoiceID     uuid.UUID
    AmountPaid    decimal.Decimal
    CurrencyPaid  Currency
    PaymentDate   time.Time
    PaymentMethod string // cash, mobile_money, bank_transfer
    Reference     string
    ExchangeRate  decimal.NullDecimal
    CreatedAt     time.Time
}