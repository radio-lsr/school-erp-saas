package repository

import (
    "context"
    "errors"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/academic"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresSectionRepository struct {
    db *pgxpool.Pool
}

func NewPostgresSectionRepository(db *pgxpool.Pool) ports.SectionRepository {
    return &PostgresSectionRepository{db: db}
}

func (r *PostgresSectionRepository) Create(ctx context.Context, section *academic.Section) error {
    query := `INSERT INTO sections (id, tenant_id, grade_level_id, academic_year_id, name, capacity, homeroom_teacher_id)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
    _, err := r.db.Exec(ctx, query,
        section.ID, section.TenantID, section.GradeLevelID, section.AcademicYearID,
        section.Name, section.Capacity, section.HomeroomTeacherID,
    )
    return err
}

func (r *PostgresSectionRepository) GetByID(ctx context.Context, id uuid.UUID) (*academic.Section, error) {
    query := `SELECT id, tenant_id, grade_level_id, academic_year_id, name, capacity, homeroom_teacher_id, created_at, updated_at
              FROM sections WHERE id = $1`
    row := r.db.QueryRow(ctx, query, id)
    var s academic.Section
    err := row.Scan(&s.ID, &s.TenantID, &s.GradeLevelID, &s.AcademicYearID, &s.Name, &s.Capacity, &s.HomeroomTeacherID, &s.CreatedAt, &s.UpdatedAt)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    return &s, nil
}

func (r *PostgresSectionRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*academic.Section, error) {
    query := `SELECT id, tenant_id, grade_level_id, academic_year_id, name, capacity, homeroom_teacher_id, created_at, updated_at
              FROM sections WHERE tenant_id = $1 ORDER BY name`
    rows, err := r.db.Query(ctx, query, tenantID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var sections []*academic.Section
    for rows.Next() {
        var s academic.Section
        err := rows.Scan(&s.ID, &s.TenantID, &s.GradeLevelID, &s.AcademicYearID, &s.Name, &s.Capacity, &s.HomeroomTeacherID, &s.CreatedAt, &s.UpdatedAt)
        if err != nil {
            return nil, err
        }
        sections = append(sections, &s)
    }
    return sections, nil
}

func (r *PostgresSectionRepository) Update(ctx context.Context, section *academic.Section) error {
    query := `UPDATE sections SET name = $1, capacity = $2, homeroom_teacher_id = $3, updated_at = NOW() WHERE id = $4`
    _, err := r.db.Exec(ctx, query, section.Name, section.Capacity, section.HomeroomTeacherID, section.ID)
    return err
}

func (r *PostgresSectionRepository) Delete(ctx context.Context, id uuid.UUID) error {
    query := `DELETE FROM sections WHERE id = $1`
    _, err := r.db.Exec(ctx, query, id)
    return err
}