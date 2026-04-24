package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radio-lsr/school-erp-saas/backend/internal/adapters/repository"
	"github.com/radio-lsr/school-erp-saas/backend/internal/config"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/ports" // <-- AJOUT
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type Application struct {
	Config            *config.Config
	DB                *pgxpool.Pool
	UserRepo          ports.UserRepository // ajouter
	SectionService    *services.SectionService
	StudentService    *services.StudentService
	EnrollmentService *services.EnrollmentService
	PaymentService    *services.PaymentService
	InvoiceGenService *services.InvoiceGenerationService
	InvoiceService    *services.InvoiceService
}

func NewApplication(db *pgxpool.Pool, cfg *config.Config) *Application {
	userRepo := repository.NewPostgresUserRepository(db)
	sectionRepo := repository.NewPostgresSectionRepository(db)
	studentRepo := repository.NewPostgresStudentRepository(db)
	enrollmentRepo := repository.NewPostgresEnrollmentRepository(db)
	feeStructRepo := repository.NewPostgresFeeStructureRepository(db)
	feeInstallmentRepo := repository.NewPostgresFeeInstallmentRepository(db)
	invoiceRepo := repository.NewPostgresInvoiceRepository(db)
	paymentRepo := repository.NewPostgresPaymentRepository(db)
	exchangeRateRepo := repository.NewPostgresExchangeRateRepository(db)

	// On ignore pour l'instant les variables non utilisées
	_, _ = feeStructRepo, feeInstallmentRepo

	sectionService := services.NewSectionService(sectionRepo)
	studentService := services.NewStudentService(studentRepo)
	enrollmentService := services.NewEnrollmentService(enrollmentRepo, sectionRepo, studentRepo)
	paymentService := services.NewPaymentService(invoiceRepo, paymentRepo, exchangeRateRepo)
	invoiceGenService := services.NewInvoiceGenerationService() // pas encore de dépendances
	invoiceService := services.NewInvoiceService(invoiceRepo, feeInstallmentRepo)

	return &Application{
		Config:            cfg,
		DB:                db,
		UserRepo:          userRepo,
		SectionService:    sectionService,
		StudentService:    studentService,
		EnrollmentService: enrollmentService,
		PaymentService:    paymentService,
		InvoiceGenService: invoiceGenService,
		InvoiceService:    invoiceService,
	}
}
