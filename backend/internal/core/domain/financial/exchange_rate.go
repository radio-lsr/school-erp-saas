package financial

import (
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

type ExchangeRate struct {
    ID            uuid.UUID
    TenantID      uuid.UUID
    FromCurrency  Currency
    ToCurrency    Currency
    Rate          decimal.Decimal
    EffectiveDate time.Time
    CreatedAt     time.Time
}