package enrollment

import (
    "time"
    "github.com/google/uuid"
)

type Enrollment struct {
    ID             uuid.UUID
    TenantID       uuid.UUID
    StudentID      uuid.UUID
    SectionID      uuid.UUID
    EnrollmentDate time.Time
    Status         string // active, transferred, graduated, dropped
    CreatedAt      time.Time
    UpdatedAt      time.Time
}