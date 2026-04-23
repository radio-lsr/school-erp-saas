package academic

import (
    "time"
    "github.com/google/uuid"
)

type GradeLevel struct {
    ID           uuid.UUID
    TenantID     uuid.UUID
    Name         string
    Cycle        string // maternelle, primaire, secondaire
    DisplayOrder int
    CreatedAt    time.Time
    UpdatedAt    time.Time
}