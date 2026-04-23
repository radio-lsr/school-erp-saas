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
    student, err := s.studentRepo.GetByID(ctx, cmd.StudentID)
    if err != nil || student == nil || student.TenantID != cmd.TenantID {
        return nil, errors.New("student not found")
    }
    section, err := s.sectionRepo.GetByID(ctx, cmd.SectionID)
    if err != nil || section == nil || section.TenantID != cmd.TenantID {
        return nil, errors.New("section not found")
    }

    activeCount, err := s.enrollmentRepo.CountActiveBySection(ctx, cmd.SectionID)
    if err != nil {
        return nil, err
    }
    if activeCount >= section.Capacity {
        return nil, errors.New("section is full")
    }

    existing, err := s.enrollmentRepo.ListByStudent(ctx, cmd.StudentID)
    if err != nil {
        return nil, err
    }
    for _, e := range existing {
        if e.SectionID == cmd.SectionID && e.Status == "active" {
            return nil, errors.New("student already enrolled in this section")
        }
    }

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