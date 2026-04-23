package services

import (
    "context"
    "errors"
    "time"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/student"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type StudentService struct {
    repo ports.StudentRepository
}

func NewStudentService(repo ports.StudentRepository) *StudentService {
    return &StudentService{repo: repo}
}

type CreateStudentCommand struct {
    TenantID  uuid.UUID
    UserID    *uuid.UUID
    FirstName string
    LastName  string
    BirthDate string // YYYY-MM-DD
    Gender    string
}

type UpdateStudentCommand struct {
    ID        uuid.UUID
    UserID    *uuid.UUID
    FirstName string
    LastName  string
    BirthDate string
    Gender    string
}

func (s *StudentService) Create(ctx context.Context, cmd CreateStudentCommand) (*student.Student, error) {
    var birthDate *time.Time
    if cmd.BirthDate != "" {
        t, err := time.Parse("2006-01-02", cmd.BirthDate)
        if err != nil {
            return nil, errors.New("invalid birth_date format (YYYY-MM-DD)")
        }
        birthDate = &t
    }

    student := &student.Student{
        ID:        uuid.New(),
        TenantID:  cmd.TenantID,
        UserID:    cmd.UserID,
        FirstName: cmd.FirstName,
        LastName:  cmd.LastName,
        BirthDate: birthDate,
        Gender:    cmd.Gender,
    }
    if err := s.repo.Create(ctx, student); err != nil {
        return nil, err
    }
    return student, nil
}

func (s *StudentService) GetByID(ctx context.Context, id uuid.UUID) (*student.Student, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *StudentService) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*student.Student, error) {
    return s.repo.ListByTenant(ctx, tenantID)
}

func (s *StudentService) Update(ctx context.Context, cmd UpdateStudentCommand) (*student.Student, error) {
    existing, err := s.repo.GetByID(ctx, cmd.ID)
    if err != nil || existing == nil {
        return nil, errors.New("student not found")
    }
    if cmd.BirthDate != "" {
        t, err := time.Parse("2006-01-02", cmd.BirthDate)
        if err != nil {
            return nil, errors.New("invalid birth_date format")
        }
        existing.BirthDate = &t
    }
    existing.UserID = cmd.UserID
    existing.FirstName = cmd.FirstName
    existing.LastName = cmd.LastName
    existing.Gender = cmd.Gender
    existing.UpdatedAt = time.Now()
    if err := s.repo.Update(ctx, existing); err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *StudentService) Delete(ctx context.Context, id uuid.UUID) error {
    return s.repo.Delete(ctx, id)
}