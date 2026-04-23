package repository

import (
    "context"
    "errors"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/student"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresStudentRepository struct {
    db *pgxpool.Pool
}

func NewPostgresStudentRepository(db *pgxpool.Pool) ports.StudentRepository {
    return &PostgresStudentRepository{db: db}
}

func (r *PostgresStudentRepository) Create(ctx context.Context, s *student.Student) error {
    query := `INSERT INTO students (id, tenant_id, user_id, first_name, last_name, birth_date, gender, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
    _, err := r.db.Exec(ctx, query,
        s.ID, s.TenantID, s.UserID, s.FirstName, s.LastName, s.BirthDate, s.Gender,
        s.CreatedAt, s.UpdatedAt)
    return err
}

func (r *PostgresStudentRepository) GetByID(ctx context.Context, id uuid.UUID) (*student.Student, error) {
    query := `SELECT id, tenant_id, user_id, first_name, last_name, birth_date, gender, created_at, updated_at
              FROM students WHERE id = $1`
    row := r.db.QueryRow(ctx, query, id)
    var s student.Student
    err := row.Scan(&s.ID, &s.TenantID, &s.UserID, &s.FirstName, &s.LastName, &s.BirthDate, &s.Gender, &s.CreatedAt, &s.UpdatedAt)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return &s, nil
}

func (r *PostgresStudentRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*student.Student, error) {
    query := `SELECT id, tenant_id, user_id, first_name, last_name, birth_date, gender, created_at, updated_at
              FROM students WHERE tenant_id = $1 ORDER BY last_name, first_name`
    rows, err := r.db.Query(ctx, query, tenantID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var students []*student.Student
    for rows.Next() {
        var s student.Student
        if err := rows.Scan(&s.ID, &s.TenantID, &s.UserID, &s.FirstName, &s.LastName, &s.BirthDate, &s.Gender, &s.CreatedAt, &s.UpdatedAt); err != nil {
            return nil, err
        }
        students = append(students, &s)
    }
    return students, nil
}

func (r *PostgresStudentRepository) Update(ctx context.Context, s *student.Student) error {
    query := `UPDATE students SET user_id = $1, first_name = $2, last_name = $3, birth_date = $4, gender = $5, updated_at = $6 WHERE id = $7`
    _, err := r.db.Exec(ctx, query, s.UserID, s.FirstName, s.LastName, s.BirthDate, s.Gender, s.UpdatedAt, s.ID)
    return err
}

func (r *PostgresStudentRepository) Delete(ctx context.Context, id uuid.UUID) error {
    query := `DELETE FROM students WHERE id = $1`
    _, err := r.db.Exec(ctx, query, id)
    return err
}