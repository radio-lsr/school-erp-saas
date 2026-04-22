func (h *PaymentHandler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		InvoiceID   string `json:"invoice_id"`
		PhoneNumber string `json:"phone_number"`
		Provider    string `json:"provider"` // "orange", "airtel", "mpesa"
	}
	// ... décoder

	// Récupérer la facture
	invoice, _ := h.invoiceService.GetByID(r.Context(), uuid.MustParse(req.InvoiceID))

	// Appeler la passerelle
	gateway := h.paymentGateway // sélection selon provider
	resp, err := gateway.InitiatePayment(r.Context(), ports.PaymentRequest{
		Amount:      invoice.TotalAmount,
		Currency:    string(invoice.Currency),
		PhoneNumber: req.PhoneNumber,
		Description: fmt.Sprintf("Frais scolaire %s", invoice.InvoiceNumber),
		Reference:   invoice.InvoiceNumber,
		CallbackURL: "https://votre-domaine.com/api/payments/callback",
	})
	// Enregistrer la tentative de paiement en base avec statut "pending"
	// ...
	json.NewEncoder(w).Encode(resp)
}