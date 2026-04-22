package repository

import (
	"context"

	"github.com/google/uuid"
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
		return nil, err
	}
	return &s, nil
}

// ... autres méthodes
