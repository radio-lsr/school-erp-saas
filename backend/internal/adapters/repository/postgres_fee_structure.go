package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresFeeStructureRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFeeStructureRepository(db *pgxpool.Pool) ports.FeeStructureRepository {
	return &PostgresFeeStructureRepository{db: db}
}

func (r *PostgresFeeStructureRepository) Create(ctx context.Context, f *financial.FeeStructure) error {
	query := `INSERT INTO fee_structures (id, tenant_id, grade_level_id, academic_year_id, name, total_amount, currency)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query, f.ID, f.TenantID, f.GradeLevelID, f.AcademicYearID, f.Name, f.TotalAmount, f.Currency)
	return err
}

func (r *PostgresFeeStructureRepository) GetByID(ctx context.Context, id uuid.UUID) (*financial.FeeStructure, error) {
	query := `SELECT id, tenant_id, grade_level_id, academic_year_id, name, total_amount, currency, created_at, updated_at
              FROM fee_structures WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	var f financial.FeeStructure
	err := row.Scan(&f.ID, &f.TenantID, &f.GradeLevelID, &f.AcademicYearID, &f.Name, &f.TotalAmount, &f.Currency, &f.CreatedAt, &f.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return &f, err
}

func (r *PostgresFeeStructureRepository) GetByGradeAndYear(ctx context.Context, tenantID, gradeLevelID, academicYearID uuid.UUID) (*financial.FeeStructure, error) {
	query := `SELECT id, tenant_id, grade_level_id, academic_year_id, name, total_amount, currency, created_at, updated_at
              FROM fee_structures WHERE tenant_id = $1 AND grade_level_id = $2 AND academic_year_id = $3`
	row := r.db.QueryRow(ctx, query, tenantID, gradeLevelID, academicYearID)
	var f financial.FeeStructure
	err := row.Scan(&f.ID, &f.TenantID, &f.GradeLevelID, &f.AcademicYearID, &f.Name, &f.TotalAmount, &f.Currency, &f.CreatedAt, &f.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return &f, err
}

func (r *PostgresFeeStructureRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*financial.FeeStructure, error) {
	query := `SELECT id, tenant_id, grade_level_id, academic_year_id, name, total_amount, currency, created_at, updated_at
              FROM fee_structures WHERE tenant_id = $1`
	rows, err := r.db.Query(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*financial.FeeStructure
	for rows.Next() {
		var f financial.FeeStructure
		if err := rows.Scan(&f.ID, &f.TenantID, &f.GradeLevelID, &f.AcademicYearID, &f.Name, &f.TotalAmount, &f.Currency, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, &f)
	}
	return list, nil
}

func (r *PostgresFeeStructureRepository) Update(ctx context.Context, f *financial.FeeStructure) error {
	query := `UPDATE fee_structures SET name = $1, total_amount = $2, currency = $3, updated_at = NOW() WHERE id = $4`
	_, err := r.db.Exec(ctx, query, f.Name, f.TotalAmount, f.Currency, f.ID)
	return err
}

func (r *PostgresFeeStructureRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM fee_structures WHERE id = $1`, id)
	return err
}
