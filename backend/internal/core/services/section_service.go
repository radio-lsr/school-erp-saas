package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/academic"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type SectionService struct {
	sectionRepo ports.SectionRepository
}

func NewSectionService(sectionRepo ports.SectionRepository) *SectionService {
	return &SectionService{sectionRepo: sectionRepo}
}

type CreateSectionCommand struct {
	TenantID          uuid.UUID
	GradeLevelID      uuid.UUID
	AcademicYearID    uuid.UUID
	Name              string
	Capacity          int
	HomeroomTeacherID *uuid.UUID
}

func (s *SectionService) CreateSection(ctx context.Context, cmd CreateSectionCommand) (*academic.Section, error) {
	section := &academic.Section{
		ID:                uuid.New(),
		TenantID:          cmd.TenantID,
		GradeLevelID:      cmd.GradeLevelID,
		AcademicYearID:    cmd.AcademicYearID,
		Name:              cmd.Name,
		Capacity:          cmd.Capacity,
		HomeroomTeacherID: cmd.HomeroomTeacherID,
	}
	if err := s.sectionRepo.Create(ctx, section); err != nil {
		return nil, err
	}
	return section, nil
}
