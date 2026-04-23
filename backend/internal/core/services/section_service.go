package services

import (
    "context"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/academic"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
)

type SectionService struct {
    repo ports.SectionRepository
}

func NewSectionService(repo ports.SectionRepository) *SectionService {
    return &SectionService{repo: repo}
}

type CreateSectionCommand struct {
    TenantID         uuid.UUID
    GradeLevelID     uuid.UUID
    AcademicYearID   uuid.UUID
    Name             string
    Capacity         int
    HomeroomTeacherID *uuid.UUID
}

func (s *SectionService) CreateSection(ctx context.Context, cmd CreateSectionCommand) (*academic.Section, error) {
    section := &academic.Section{
        ID:               uuid.New(),
        TenantID:         cmd.TenantID,
        GradeLevelID:     cmd.GradeLevelID,
        AcademicYearID:   cmd.AcademicYearID,
        Name:             cmd.Name,
        Capacity:         cmd.Capacity,
        HomeroomTeacherID: cmd.HomeroomTeacherID,
    }
    if err := s.repo.Create(ctx, section); err != nil {
        return nil, err
    }
    return section, nil
}

func (s *SectionService) GetByID(ctx context.Context, id uuid.UUID) (*academic.Section, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *SectionService) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*academic.Section, error) {
    return s.repo.ListByTenant(ctx, tenantID)
}

func (s *SectionService) UpdateSection(ctx context.Context, section *academic.Section) error {
    return s.repo.Update(ctx, section)
}

func (s *SectionService) DeleteSection(ctx context.Context, id uuid.UUID) error {
    return s.repo.Delete(ctx, id)
}