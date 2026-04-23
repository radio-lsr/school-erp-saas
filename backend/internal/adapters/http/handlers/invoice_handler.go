package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type InvoiceHandler struct {
    invoiceGenService *services.InvoiceGenerationService
}

func NewInvoiceHandler(service *services.InvoiceGenerationService) *InvoiceHandler {
    return &InvoiceHandler{invoiceGenService: service}
}

func (h *InvoiceHandler) GenerateInvoices(w http.ResponseWriter, r *http.Request) {
    var req struct {
        AcademicYearID string `json:"academic_year_id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid body", http.StatusBadRequest)
        return
    }
    academicYearID, _ := uuid.Parse(req.AcademicYearID)
    // TenantID from context? simplified
    err := h.invoiceGenService.GenerateInvoicesForAcademicYear(r.Context(), uuid.Nil, academicYearID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{"message": "generation started"})
}