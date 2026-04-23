package repository

import (
    "context"
    "errors"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/enrollment"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type PostgresEnrollmentRepository struct {
    db *pgxpool.Pool
}

func NewPostgresEnrollmentRepository(db *pgxpool.Pool) ports.EnrollmentRepository {
    return &PostgresEnrollmentRepository{db: db}
}

func (r *PostgresEnrollmentRepository) Create(ctx context.Context, e *enrollment.Enrollment) error {
    query := `INSERT INTO enrollments (id, tenant_id, student_id, section_id, enrollment_date, status, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
    _, err := r.db.Exec(ctx, query, e.ID, e.TenantID, e.StudentID, e.SectionID, e.EnrollmentDate, e.Status, e.CreatedAt, e.UpdatedAt)
    return err
}

func (r *PostgresEnrollmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*enrollment.Enrollment, error) {
    query := `SELECT id, tenant_id, student_id, section_id, enrollment_date, status, created_at, updated_at FROM enrollments WHERE id = $1`
    row := r.db.QueryRow(ctx, query, id)
    var e enrollment.Enrollment
    err := row.Scan(&e.ID, &e.TenantID, &e.StudentID, &e.SectionID, &e.EnrollmentDate, &e.Status, &e.CreatedAt, &e.UpdatedAt)
    if errors.Is(err, pgx.ErrNoRows) {
        return nil, nil
    }
    return &e, err
}

func (r *PostgresEnrollmentRepository) ListByStudent(ctx context.Context, studentID uuid.UUID) ([]*enrollment.Enrollment, error) {
    query := `SELECT id, tenant_id, student_id, section_id, enrollment_date, status, created_at, updated_at FROM enrollments WHERE student_id = $1`
    rows, err := r.db.Query(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var enrollments []*enrollment.Enrollment
    for rows.Next() {
        var e enrollment.Enrollment
        if err := rows.Scan(&e.ID, &e.TenantID, &e.StudentID, &e.SectionID, &e.EnrollmentDate, &e.Status, &e.CreatedAt, &e.UpdatedAt); err != nil {
            return nil, err
        }
        enrollments = append(enrollments, &e)
    }
    return enrollments, nil
}

func (r *PostgresEnrollmentRepository) ListBySection(ctx context.Context, sectionID uuid.UUID) ([]*enrollment.Enrollment, error) {
    query := `SELECT id, tenant_id, student_id, section_id, enrollment_date, status, created_at, updated_at FROM enrollments WHERE section_id = $1`
    rows, err := r.db.Query(ctx, query, sectionID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var enrollments []*enrollment.Enrollment
    for rows.Next() {
        var e enrollment.Enrollment
        if err := rows.Scan(&e.ID, &e.TenantID, &e.StudentID, &e.SectionID, &e.EnrollmentDate, &e.Status, &e.CreatedAt, &e.UpdatedAt); err != nil {
            return nil, err
        }
        enrollments = append(enrollments, &e)
    }
    return enrollments, nil
}

func (r *PostgresEnrollmentRepository) CountActiveBySection(ctx context.Context, sectionID uuid.UUID) (int, error) {
    var count int
    err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM enrollments WHERE section_id = $1 AND status = 'active'`, sectionID).Scan(&count)
    return count, err
}

func (r *PostgresEnrollmentRepository) Update(ctx context.Context, e *enrollment.Enrollment) error {
    query := `UPDATE enrollments SET status = $1, updated_at = $2 WHERE id = $3`
    _, err := r.db.Exec(ctx, query, e.Status, e.UpdatedAt, e.ID)
    return err
}

func (r *PostgresEnrollmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
    _, err := r.db.Exec(ctx, `DELETE FROM enrollments WHERE id = $1`, id)
    return err
}

func (r *PostgresEnrollmentRepository) ListActiveByAcademicYear(ctx context.Context, tenantID, academicYearID uuid.UUID) ([]*enrollment.Enrollment, error) {
    query := `SELECT e.id, e.tenant_id, e.student_id, e.section_id, e.enrollment_date, e.status, e.created_at, e.updated_at
              FROM enrollments e
              JOIN sections s ON e.section_id = s.id
              WHERE e.tenant_id = $1 AND s.academic_year_id = $2 AND e.status = 'active'`
    rows, err := r.db.Query(ctx, query, tenantID, academicYearID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var enrollments []*enrollment.Enrollment
    for rows.Next() {
        var e enrollment.Enrollment
        if err := rows.Scan(&e.ID, &e.TenantID, &e.StudentID, &e.SectionID, &e.EnrollmentDate, &e.Status, &e.CreatedAt, &e.UpdatedAt); err != nil {
            return nil, err
        }
        enrollments = append(enrollments, &e)
    }
    return enrollments, nil
}