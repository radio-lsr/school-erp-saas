package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/radio-lsr/school-erp-saas/backend/internal/app"
)

type PaymentHandler struct {
    app *app.Application
}

func NewPaymentHandler(app *app.Application) *PaymentHandler {
    return &PaymentHandler{app: app}
}

// InitiatePayment démarre un paiement Mobile Money
func (h *PaymentHandler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Payment initiation not implemented yet",
    })
}

// Callback reçoit les notifications de la passerelle de paiement
func (h *PaymentHandler) Callback(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}