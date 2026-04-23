package academic

import (
    "time"
    "github.com/google/uuid"
)

type AcademicYear struct {
    ID        uuid.UUID
    TenantID  uuid.UUID
    Name      string
    StartDate time.Time
    EndDate   time.Time
    IsCurrent bool
    CreatedAt time.Time
    UpdatedAt time.Time
}