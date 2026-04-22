package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/radio-lsr/school-erp-saas/backend/internal/core/domain/financial"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports"
	"github.com/shopspring/decimal"
)

type PaymentCallbackHandler struct {
	paymentService *services.PaymentService // service pour ajouter le paiement
	invoiceRepo    ports.InvoiceRepository  // pour récupérer la facture
	paymentRepo    ports.PaymentRepository  // pour vérifier l'idempotence
	gateway        ports.PaymentGateway     // pour vérifier le statut si nécessaire (optionnel)
}

func NewPaymentCallbackHandler(
	paymentService *services.PaymentService,
	invoiceRepo ports.InvoiceRepository,
	paymentRepo ports.PaymentRepository,
) *PaymentCallbackHandler {
	return &PaymentCallbackHandler{
		paymentService: paymentService,
		invoiceRepo:    invoiceRepo,
		paymentRepo:    paymentRepo,
	}
}

// Structure attendue de FlexPay (à adapter selon leur documentation réelle)
type FlexPayCallbackPayload struct {
	TransactionID string `json:"transaction_id"`
	Reference     string `json:"reference"` // numéro de facture
	Status        string `json:"status"`    // "success", "failed", "pending"
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Phone         string `json:"phone"`
	Signature     string `json:"signature"` // pour vérifier l'authenticité
}

func (h *PaymentCallbackHandler) FlexPayCallback(w http.ResponseWriter, r *http.Request) {
	var payload FlexPayCallbackPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// 1. Vérifier la signature (à implémenter avec votre clé secrète FlexPay)
	// if !verifyFlexPaySignature(payload, r.Header.Get("X-FlexPay-Signature")) {
	//     http.Error(w, "invalid signature", http.StatusUnauthorized)
	//     return
	// }

	// 2. Vérifier si ce paiement n'a pas déjà été traité (idempotence)
	existing, _ := h.paymentRepo.GetByTransactionID(r.Context(), payload.TransactionID)
	if existing != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("already processed"))
		return
	}

	// 3. Si le statut n'est pas "success", on peut logger et ignorer
	if payload.Status != "success" {
		// Log l'échec, mais répondre OK pour éviter les retries
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("status not success"))
		return
	}

	// 4. Récupérer la facture via le numéro de facture (Reference)
	invoice, err := h.invoiceRepo.GetByNumber(r.Context(), payload.Reference)
	if err != nil {
		http.Error(w, "invoice not found", http.StatusNotFound)
		return
	}

	// 5. Créer le paiement via le service
	amount, _ := decimal.NewFromString(payload.Amount)
	cmd := services.AddPaymentCommand{
		InvoiceID:     invoice.ID,
		AmountPaid:    amount,
		CurrencyPaid:  financial.Currency(payload.Currency),
		PaymentMethod: "mobile_money",
		Reference:     payload.TransactionID,
	}

	_, err = h.paymentService.AddPayment(r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
