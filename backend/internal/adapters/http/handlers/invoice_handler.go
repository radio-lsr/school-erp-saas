package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/radio-lsr/school-erp-saas/backend/internal/adapters/http/middleware"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type InvoiceHandler struct {
	invoiceGenService *services.InvoiceGenerationService
	// ... autres services si besoin
}

func NewInvoiceHandler(invoiceGenService *services.InvoiceGenerationService) *InvoiceHandler {
	return &InvoiceHandler{invoiceGenService: invoiceGenService}
}

type generateInvoicesRequest struct {
	AcademicYearID string `json:"academic_year_id"`
}

// GenerateInvoices déclenche la génération des factures pour une année académique.
// Accessible uniquement aux administrateurs (vérifié par le middleware).
func (h *InvoiceHandler) GenerateInvoices(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Context().Value(middleware.TenantIDKey).(uuid.UUID)

	var req generateInvoicesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	academicYearID, err := uuid.Parse(req.AcademicYearID)
	if err != nil {
		http.Error(w, "invalid academic_year_id", http.StatusBadRequest)
		return
	}

	err = h.invoiceGenService.GenerateInvoicesForAcademicYear(r.Context(), tenantID, academicYearID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "invoice generation started"})
}

func (h *InvoiceHandler) Routes() chi.Router {
	r := chi.NewRouter()
	// Ajouter d'autres routes pour les factures si nécessaire
	return r
}
