package ports

import (
    "context"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/user"
)

type UserRepository interface {
    GetByEmail(ctx context.Context, email string) (*user.User, error)
}