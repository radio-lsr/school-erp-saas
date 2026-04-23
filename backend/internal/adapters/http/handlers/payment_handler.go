package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
    "github.com/shopspring/decimal"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
)

type PaymentHandler struct {
    service *services.PaymentService
}

func NewPaymentHandler(service *services.PaymentService) *PaymentHandler {
    return &PaymentHandler{service: service}
}

func (h *PaymentHandler) AddPayment(w http.ResponseWriter, r *http.Request) {
    var req struct {
        InvoiceID     string `json:"invoice_id"`
        Amount        string `json:"amount"`
        Currency      string `json:"currency"`
        PaymentMethod string `json:"payment_method"`
        Reference     string `json:"reference"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    invoiceID, _ := uuid.Parse(req.InvoiceID)
    amount, _ := decimal.NewFromString(req.Amount)
    cmd := services.AddPaymentCommand{
        InvoiceID:     invoiceID,
        AmountPaid:    amount,
        CurrencyPaid:  financial.Currency(req.Currency),
        PaymentMethod: req.PaymentMethod,
        Reference:     req.Reference,
    }
    payment, err := h.service.AddPayment(r.Context(), cmd)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(payment)
}

func (h *PaymentHandler) Routes() chi.Router {
    r := chi.NewRouter()
    r.Post("/", h.AddPayment)
    return r
}