package academic

import (
    "time"
    "github.com/google/uuid"
)

type Section struct {
    ID               uuid.UUID
    TenantID         uuid.UUID
    GradeLevelID     uuid.UUID
    AcademicYearID   uuid.UUID
    Name             string
    Capacity         int
    HomeroomTeacherID *uuid.UUID
    CreatedAt        time.Time
    UpdatedAt        time.Time
}