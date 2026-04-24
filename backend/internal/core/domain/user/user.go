package user

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    ID           uuid.UUID
    TenantID     uuid.UUID
    Email        string
    PasswordHash string
    FullName     string
    Role         string
    CreatedAt    time.Time
    UpdatedAt    time.Time
}