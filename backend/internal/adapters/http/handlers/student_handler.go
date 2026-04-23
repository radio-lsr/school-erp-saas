package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/adapters/http/middleware"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type StudentHandler struct {
    studentService *services.StudentService
}

func NewStudentHandler(studentService *services.StudentService) *StudentHandler {
    return &StudentHandler{studentService: studentService}
}

type createStudentRequest struct {
    UserID    *string `json:"user_id"`
    FirstName string  `json:"first_name"`
    LastName  string  `json:"last_name"`
    BirthDate string  `json:"birth_date"` // YYYY-MM-DD
    Gender    string  `json:"gender"`
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
    tenantID := r.Context().Value(middleware.TenantIDKey).(uuid.UUID)

    var req createStudentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    var userID *uuid.UUID
    if req.UserID != nil && *req.UserID != "" {
        id := uuid.MustParse(*req.UserID)
        userID = &id
    }

    cmd := services.CreateStudentCommand{
        TenantID:  tenantID,
        UserID:    userID,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        BirthDate: req.BirthDate,  // passage direct de la chaîne
        Gender:    req.Gender,
    }

    student, err := h.studentService.Create(r.Context(), cmd)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    student, err := h.studentService.GetByID(r.Context(), id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if student == nil {
        http.Error(w, "student not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
    tenantID := r.Context().Value(middleware.TenantIDKey).(uuid.UUID)

    students, err := h.studentService.ListByTenant(r.Context(), tenantID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(students)
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    var req createStudentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    var userID *uuid.UUID
    if req.UserID != nil && *req.UserID != "" {
        uid := uuid.MustParse(*req.UserID)
        userID = &uid
    }

    cmd := services.UpdateStudentCommand{
        ID:        id,
        UserID:    userID,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        BirthDate: req.BirthDate,
        Gender:    req.Gender,
    }

    student, err := h.studentService.Update(r.Context(), cmd)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    if err := h.studentService.Delete(r.Context(), id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (h *StudentHandler) Routes() chi.Router {
    r := chi.NewRouter()
    r.Post("/", h.Create)
    r.Get("/", h.List)
    r.Get("/{id}", h.GetByID)
    r.Put("/{id}", h.Update)
    r.Delete("/{id}", h.Delete)
    return r
}