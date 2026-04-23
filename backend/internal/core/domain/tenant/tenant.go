package tenant

import (
    "time"
    "github.com/google/uuid"
)

type Tenant struct {
    ID        uuid.UUID
    Name      string
    Subdomain string
    CreatedAt time.Time
    UpdatedAt time.Time
}