package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "github.com/go-chi/chi/v5"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/adapters/http/middleware"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type EnrollmentHandler struct {
    service *services.EnrollmentService
}

func NewEnrollmentHandler(service *services.EnrollmentService) *EnrollmentHandler {
    return &EnrollmentHandler{service: service}
}

func (h *EnrollmentHandler) Enroll(w http.ResponseWriter, r *http.Request) {
    tenantID := r.Context().Value(middleware.TenantIDKey).(uuid.UUID)
    var req struct {
        StudentID      string `json:"student_id"`
        SectionID      string `json:"section_id"`
        EnrollmentDate string `json:"enrollment_date"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    studentID, _ := uuid.Parse(req.StudentID)
    sectionID, _ := uuid.Parse(req.SectionID)
    enrollmentDate := time.Now()
    if req.EnrollmentDate != "" {
        enrollmentDate, _ = time.Parse("2006-01-02", req.EnrollmentDate)
    }

    cmd := services.EnrollStudentCommand{
        TenantID:       tenantID,
        StudentID:      studentID,
        SectionID:      sectionID,
        EnrollmentDate: enrollmentDate,
    }
    enrollment, err := h.service.EnrollStudent(r.Context(), cmd)
    if err != nil {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(enrollment)
}

func (h *EnrollmentHandler) Routes() chi.Router {
    r := chi.NewRouter()
    r.Post("/", h.Enroll)
    return r
}