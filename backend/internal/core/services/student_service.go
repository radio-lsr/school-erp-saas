package services

import (
	"context"

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
	BirthDate string // format YYYY-MM-DD
	Gender    string
}

func (s *StudentService) Create(ctx context.Context, cmd CreateStudentCommand) (*student.Student, error) {
	// parsing birthDate, validation...
	student := &student.Student{
		ID:        uuid.New(),
		TenantID:  cmd.TenantID,
		UserID:    cmd.UserID,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Gender:    cmd.Gender,
	}
	if err := s.repo.Create(ctx, student); err != nil {
		return nil, err
	}
	return student, nil
}
