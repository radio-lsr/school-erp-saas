package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/radio-lsr/school-erp-saas/backend/internal/adapters/http/middleware"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type InvoiceHandler struct {
	invoiceGenService *services.InvoiceGenerationService
	// on ajoute un service de création simple (ou on élargit le service existant)
	invoiceService *services.InvoiceService // à créer si nécessaire, ou utiliser un repo directement
}

func NewInvoiceHandler(genService *services.InvoiceGenerationService, invService *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceGenService: genService,
		invoiceService:    invService,
	}
}

// CreateInvoice crée une facture pour un étudiant, une échéance donnée.
func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Context().Value(middleware.TenantIDKey).(uuid.UUID)
	var req struct {
		StudentID        string `json:"student_id"`
		FeeInstallmentID string `json:"fee_installment_id"`
		Amount           string `json:"amount"`
		Currency         string `json:"currency"`
		DueDate          string `json:"due_date"` // optionnel, sinon on utilise la date de l'échéance
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	studentID, _ := uuid.Parse(req.StudentID)
	installmentID, _ := uuid.Parse(req.FeeInstallmentID)

	invoice, err := h.invoiceService.CreateInvoice(r.Context(), services.CreateInvoiceCommand{
		TenantID:         tenantID,
		StudentID:        studentID,
		FeeInstallmentID: installmentID,
		Amount:           req.Amount,
		Currency:         financial.Currency(req.Currency),
		DueDate:          req.DueDate,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

// GenerateInvoices (déjà présente)
func (h *InvoiceHandler) GenerateInvoices(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AcademicYearID string `json:"academic_year_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	academicYearID, _ := uuid.Parse(req.AcademicYearID)
	tenantID := r.Context().Value(middleware.TenantIDKey).(uuid.UUID)
	err := h.invoiceGenService.GenerateInvoicesForAcademicYear(r.Context(), tenantID, academicYearID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "generation started"})
}
