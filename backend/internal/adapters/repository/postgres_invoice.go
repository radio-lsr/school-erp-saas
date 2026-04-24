package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresInvoiceRepository struct {
	db *pgxpool.Pool
}

func NewPostgresInvoiceRepository(db *pgxpool.Pool) ports.InvoiceRepository {
	return &PostgresInvoiceRepository{db: db}
}

func (r *PostgresInvoiceRepository) Create(ctx context.Context, invoice *financial.Invoice) error {
	query := `INSERT INTO invoices (id, tenant_id, student_id, fee_installment_id, invoice_number, total_amount, currency, status, issued_date, due_date, created_at, updated_at)
              VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err := r.db.Exec(ctx, query,
		invoice.ID, invoice.TenantID, invoice.StudentID, invoice.FeeInstallmentID,
		invoice.InvoiceNumber, invoice.TotalAmount, invoice.Currency,
		invoice.Status, invoice.IssuedDate, invoice.DueDate,
		invoice.CreatedAt, invoice.UpdatedAt)
	return err
}

func (r *PostgresInvoiceRepository) GetByID(ctx context.Context, id uuid.UUID) (*financial.Invoice, error) {
	query := `SELECT id, tenant_id, student_id, fee_installment_id, invoice_number, total_amount, currency, status, issued_date, due_date, created_at, updated_at
              FROM invoices WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	var inv financial.Invoice
	err := row.Scan(&inv.ID, &inv.TenantID, &inv.StudentID, &inv.FeeInstallmentID,
		&inv.InvoiceNumber, &inv.TotalAmount, &inv.Currency, &inv.Status,
		&inv.IssuedDate, &inv.DueDate, &inv.CreatedAt, &inv.UpdatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &inv, nil
}

func (r *PostgresInvoiceRepository) GetByNumber(ctx context.Context, number string) (*financial.Invoice, error) {
	query := `SELECT id, tenant_id, student_id, fee_installment_id, invoice_number, total_amount, currency, status, issued_date, due_date, created_at, updated_at
              FROM invoices WHERE invoice_number = $1`
	row := r.db.QueryRow(ctx, query, number)
	var inv financial.Invoice
	err := row.Scan(&inv.ID, &inv.TenantID, &inv.StudentID, &inv.FeeInstallmentID,
		&inv.InvoiceNumber, &inv.TotalAmount, &inv.Currency, &inv.Status,
		&inv.IssuedDate, &inv.DueDate, &inv.CreatedAt, &inv.UpdatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &inv, nil
}

func (r *PostgresInvoiceRepository) ListByStudent(ctx context.Context, studentID uuid.UUID) ([]*financial.Invoice, error) {
	query := `SELECT id, tenant_id, student_id, fee_installment_id, invoice_number, total_amount, currency, status, issued_date, due_date, created_at, updated_at
              FROM invoices WHERE student_id = $1 ORDER BY due_date`
	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*financial.Invoice
	for rows.Next() {
		var inv financial.Invoice
		if err := rows.Scan(&inv.ID, &inv.TenantID, &inv.StudentID, &inv.FeeInstallmentID,
			&inv.InvoiceNumber, &inv.TotalAmount, &inv.Currency, &inv.Status,
			&inv.IssuedDate, &inv.DueDate, &inv.CreatedAt, &inv.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, &inv)
	}
	return list, nil
}

func (r *PostgresInvoiceRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*financial.Invoice, error) {
	query := `SELECT id, tenant_id, student_id, fee_installment_id, invoice_number, total_amount, currency, status, issued_date, due_date, created_at, updated_at
              FROM invoices WHERE tenant_id = $1 ORDER BY due_date`
	rows, err := r.db.Query(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*financial.Invoice
	for rows.Next() {
		var inv financial.Invoice
		if err := rows.Scan(&inv.ID, &inv.TenantID, &inv.StudentID, &inv.FeeInstallmentID,
			&inv.InvoiceNumber, &inv.TotalAmount, &inv.Currency, &inv.Status,
			&inv.IssuedDate, &inv.DueDate, &inv.CreatedAt, &inv.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, &inv)
	}
	return list, nil
}

func (r *PostgresInvoiceRepository) Update(ctx context.Context, invoice *financial.Invoice) error {
	query := `UPDATE invoices SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(ctx, query, invoice.Status, invoice.ID)
	return err
}
