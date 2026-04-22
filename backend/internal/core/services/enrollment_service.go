package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/enrollment"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type EnrollmentService struct {
	enrollmentRepo ports.EnrollmentRepository
	sectionRepo    ports.SectionRepository
	studentRepo    ports.StudentRepository
}

func NewEnrollmentService(
	enrollmentRepo ports.EnrollmentRepository,
	sectionRepo ports.SectionRepository,
	studentRepo ports.StudentRepository,
) *EnrollmentService {
	return &EnrollmentService{
		enrollmentRepo: enrollmentRepo,
		sectionRepo:    sectionRepo,
		studentRepo:    studentRepo,
	}
}

type EnrollStudentCommand struct {
	TenantID       uuid.UUID
	StudentID      uuid.UUID
	SectionID      uuid.UUID
	EnrollmentDate time.Time
}

func (s *EnrollmentService) EnrollStudent(ctx context.Context, cmd EnrollStudentCommand) (*enrollment.Enrollment, error) {
	// 1. Vérifier que l'étudiant existe et appartient au tenant
	student, err := s.studentRepo.GetByID(ctx, cmd.StudentID)
	if err != nil {
		return nil, err
	}
	if student == nil || student.TenantID != cmd.TenantID {
		return nil, errors.New("student not found")
	}

	// 2. Vérifier que la section existe et appartient au tenant
	section, err := s.sectionRepo.GetByID(ctx, cmd.SectionID)
	if err != nil {
		return nil, err
	}
	if section == nil || section.TenantID != cmd.TenantID {
		return nil, errors.New("section not found")
	}

	// 3. Vérifier la capacité de la section
	activeCount, err := s.enrollmentRepo.CountActiveBySection(ctx, cmd.SectionID)
	if err != nil {
		return nil, err
	}
	if activeCount >= section.Capacity {
		return nil, errors.New("section is full")
	}

	// 4. Vérifier que l'étudiant n'est pas déjà inscrit dans cette section
	existing, err := s.enrollmentRepo.ListByStudent(ctx, cmd.StudentID)
	if err != nil {
		return nil, err
	}
	for _, e := range existing {
		if e.SectionID == cmd.SectionID && e.Status == "active" {
			return nil, errors.New("student already enrolled in this section")
		}
	}

	// 5. Créer l'inscription
	enrollmentEntity := &enrollment.Enrollment{
		ID:             uuid.New(),
		TenantID:       cmd.TenantID,
		StudentID:      cmd.StudentID,
		SectionID:      cmd.SectionID,
		EnrollmentDate: cmd.EnrollmentDate,
		Status:         "active",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.enrollmentRepo.Create(ctx, enrollmentEntity); err != nil {
		return nil, err
	}

	return enrollmentEntity, nil
}
