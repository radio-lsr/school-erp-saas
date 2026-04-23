package student

import (
    "time"
    "github.com/google/uuid"
)

type Student struct {
    ID        uuid.UUID
    TenantID  uuid.UUID
    UserID    *uuid.UUID
    FirstName string
    LastName  string
    BirthDate *time.Time
    Gender    string
    CreatedAt time.Time
    UpdatedAt time.Time
}