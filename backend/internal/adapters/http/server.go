package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/radio-lsr/school-erp-saas/backend/internal/adapters/http/handlers"
	"github.com/radio-lsr/school-erp-saas/backend/internal/adapters/http/middleware"
	"github.com/radio-lsr/school-erp-saas/backend/internal/app"
	"github.com/radio-lsr/school-erp-saas/backend/internal/config"
)

func NewServer(cfg *config.Config, application *app.Application) *http.Server {
	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// --- Routes publiques ---
	r.Post("/api/auth/login", handlers.NewAuthHandler(application).Login)

	// Webhook de paiement (public, sécurisé par signature)
	r.Post("/api/webhooks/flexpay", handlers.NewPaymentCallbackHandler().FlexPayCallback)

	// --- Routes protégées (JWT requis) ---
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(cfg))

		// Sections / Classes
		r.Mount("/api/sections", handlers.NewSectionHandler(application.SectionService).Routes())

		// Élèves
		r.Mount("/api/students", handlers.NewStudentHandler(application.StudentService).Routes())

		// Inscriptions
		r.Mount("/api/enrollments", handlers.NewEnrollmentHandler(application.EnrollmentService).Routes())

		// Paiements
		r.Mount("/api/payments", handlers.NewPaymentHandler(application.PaymentService).Routes())

		// Factures : un seul handler avec les deux services
		invoiceHandler := handlers.NewInvoiceHandler(application.InvoiceGenService, application.InvoiceService)
		r.Post("/api/invoices", invoiceHandler.CreateInvoice)
		r.Post("/api/invoices/generate", invoiceHandler.GenerateInvoices)
	})

	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}
}
