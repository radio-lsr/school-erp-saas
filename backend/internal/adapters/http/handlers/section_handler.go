package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
    "github.com/radio-lsr/school-erp-saas/backend/internal/adapters/http/middleware"
)

type SectionHandler struct {
    sectionService *services.SectionService
}

func NewSectionHandler(sectionService *services.SectionService) *SectionHandler {
    return &SectionHandler{sectionService: sectionService}
}

func (h *SectionHandler) Create(w http.ResponseWriter, r *http.Request) {
    tenantID := r.Context().Value(middleware.TenantIDKey).(uuid.UUID)
    var req struct {
        GradeLevelID     string  `json:"grade_level_id"`
        AcademicYearID   string  `json:"academic_year_id"`
        Name             string  `json:"name"`
        Capacity         int     `json:"capacity"`
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
        TenantID:         tenantID,
        GradeLevelID:     uuid.MustParse(req.GradeLevelID),
        AcademicYearID:   uuid.MustParse(req.AcademicYearID),
        Name:             req.Name,
        Capacity:         req.Capacity,
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

func (h *SectionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    section, err := h.sectionService.GetByID(r.Context(), id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if section == nil {
        http.Error(w, "section not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(section)
}

func (h *SectionHandler) Routes() chi.Router {
    r := chi.NewRouter()
    r.Post("/", h.Create)
    r.Get("/{id}", h.GetByID)
    // autres routes à venir
    return r
}