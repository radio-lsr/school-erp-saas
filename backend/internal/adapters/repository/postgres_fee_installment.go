package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresFeeInstallmentRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFeeInstallmentRepository(db *pgxpool.Pool) ports.FeeInstallmentRepository {
	return &PostgresFeeInstallmentRepository{db: db}
}

func (r *PostgresFeeInstallmentRepository) Create(ctx context.Context, i *financial.FeeInstallment) error {
	query := `INSERT INTO fee_installments (id, tenant_id, fee_structure_id, period_name, amount, currency, due_date)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query,
		i.ID, i.TenantID, i.FeeStructureID, i.PeriodName, i.Amount, i.Currency, i.DueDate)
	return err
}

func (r *PostgresFeeInstallmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*financial.FeeInstallment, error) {
	query := `SELECT id, tenant_id, fee_structure_id, period_name, amount, currency, due_date, created_at, updated_at
              FROM fee_installments WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	var i financial.FeeInstallment
	err := row.Scan(&i.ID, &i.TenantID, &i.FeeStructureID, &i.PeriodName, &i.Amount, &i.Currency, &i.DueDate, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &i, nil
}

func (r *PostgresFeeInstallmentRepository) ListByFeeStructure(ctx context.Context, feeStructureID uuid.UUID) ([]*financial.FeeInstallment, error) {
	query := `SELECT id, tenant_id, fee_structure_id, period_name, amount, currency, due_date, created_at, updated_at
              FROM fee_installments WHERE fee_structure_id = $1 ORDER BY due_date`
	rows, err := r.db.Query(ctx, query, feeStructureID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*financial.FeeInstallment
	for rows.Next() {
		var i financial.FeeInstallment
		if err := rows.Scan(&i.ID, &i.TenantID, &i.FeeStructureID, &i.PeriodName, &i.Amount, &i.Currency, &i.DueDate, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, &i)
	}
	return list, nil
}

func (r *PostgresFeeInstallmentRepository) Update(ctx context.Context, i *financial.FeeInstallment) error {
	query := `UPDATE fee_installments SET period_name = $1, amount = $2, currency = $3, due_date = $4, updated_at = NOW() WHERE id = $5`
	_, err := r.db.Exec(ctx, query, i.PeriodName, i.Amount, i.Currency, i.DueDate, i.ID)
	return err
}

func (r *PostgresFeeInstallmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM fee_installments WHERE id = $1`, id)
	return err
}
