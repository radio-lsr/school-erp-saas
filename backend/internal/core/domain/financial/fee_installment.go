package financial

import (
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

type FeeInstallment struct {
    ID             uuid.UUID
    TenantID       uuid.UUID
    FeeStructureID uuid.UUID
    PeriodName     string
    Amount         decimal.Decimal
    Currency       Currency
    DueDate        time.Time
    CreatedAt      time.Time
    UpdatedAt      time.Time
}