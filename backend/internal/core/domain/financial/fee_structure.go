package financial

import (
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

type FeeStructure struct {
    ID             uuid.UUID
    TenantID       uuid.UUID
    GradeLevelID   uuid.UUID
    AcademicYearID uuid.UUID
    Name           string
    TotalAmount    decimal.Decimal
    Currency       Currency
    CreatedAt      time.Time
    UpdatedAt      time.Time
}