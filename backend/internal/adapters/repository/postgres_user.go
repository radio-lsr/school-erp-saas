package repository

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/user"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
    "github.com/jackc/pgx/v5"
)

type PostgresUserRepository struct {
    db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) ports.UserRepository {
    return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
    query := `SELECT id, tenant_id, email, password_hash, full_name, role, created_at, updated_at
              FROM users WHERE email = $1`
    row := r.db.QueryRow(ctx, query, email)
    var u user.User
    err := row.Scan(&u.ID, &u.TenantID, &u.Email, &u.PasswordHash, &u.FullName, &u.Role, &u.CreatedAt, &u.UpdatedAt)
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &u, nil
}