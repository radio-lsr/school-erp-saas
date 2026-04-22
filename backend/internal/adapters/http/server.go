// Routes publiques (sans authentification)
r.Post("/api/auth/login", handlers.NewAuthHandler(application).Login)

// Webhooks de paiement (publics)
paymentCallbackHandler := handlers.NewPaymentCallbackHandler(
	application.PaymentService,
	application.InvoiceRepo,
	application.PaymentRepo,
)

r.Group(func(r chi.Router) {
	r.Use(middleware.AuthMiddleware(cfg))
	r.Use(middleware.RequireRole("admin"))
	r.Post("/api/invoices/generate", invoiceHandler.GenerateInvoices)
})
r.Post("/api/webhooks/flexpay", paymentCallbackHandler.FlexPayCallback)