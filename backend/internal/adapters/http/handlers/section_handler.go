package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type SectionHandler struct {
	sectionService *services.SectionService
}

func NewSectionHandler(sectionService *services.SectionService) *SectionHandler {
	return &SectionHandler{sectionService: sectionService}
}

func (h *SectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Context().Value("tenantID").(uuid.UUID)
	var req struct {
		GradeLevelID      string  `json:"grade_level_id"`
		AcademicYearID    string  `json:"academic_year_id"`
		Name              string  `json:"name"`
		Capacity          int     `json:"capacity"`
		HomeroomTeacherID *string `json:"homeroom_teacher_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var teacherID *uuid.UUID
	if req.HomeroomTeacherID != nil {
		id := uuid.MustParse(*req.HomeroomTeacherID)
		teacherID = &id
	}

	cmd := services.CreateSectionCommand{
		TenantID:          tenantID,
		GradeLevelID:      uuid.MustParse(req.GradeLevelID),
		AcademicYearID:    uuid.MustParse(req.AcademicYearID),
		Name:              req.Name,
		Capacity:          req.Capacity,
		HomeroomTeacherID: teacherID,
	}

	section, err := h.sectionService.CreateSection(r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(section)
}

func (h *SectionHandler) RegisterRoutes(r chi.Router) {
	r.Post("/api/sections", h.Create)
	r.Get("/api/sections/{id}", h.GetByID)
	// ...
}
