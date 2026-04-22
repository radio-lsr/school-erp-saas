package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radio-lsr/school-erp-saas/backend/internal/adapters/repository"
	"github.com/radio-lsr/school-erp-saas/backend/internal/config"
	"github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
)

type Application struct {
	Config            *config.Config
	DB                *pgxpool.Pool
	SectionService    *services.SectionService
	StudentService    *services.StudentService
	EnrollmentService *services.EnrollmentService
	InvoiceGenService *services.InvoiceGenerationService
	PaymentService    *services.PaymentService
	// Repositories exposés si nécessaire
	InvoiceRepo ports.InvoiceRepository
	PaymentRepo ports.PaymentRepository
}

func NewApplication(db *pgxpool.Pool, cfg *config.Config) *Application {
	// Repositories
	sectionRepo := repository.NewPostgresSectionRepository(db)
	studentRepo := repository.NewPostgresStudentRepository(db)
	enrollmentRepo := repository.NewPostgresEnrollmentRepository(db)
	feeStructureRepo := repository.NewPostgresFeeStructureRepository(db)
	feeInstallmentRepo := repository.NewPostgresFeeInstallmentRepository(db)
	invoiceRepo := repository.NewPostgresInvoiceRepository(db)
	paymentRepo := repository.NewPostgresPaymentRepository(db)
	exchangeRateRepo := repository.NewPostgresExchangeRateRepository(db)

	// Services
	sectionService := services.NewSectionService(sectionRepo)
	studentService := services.NewStudentService(studentRepo)
	enrollmentService := services.NewEnrollmentService(enrollmentRepo, sectionRepo, studentRepo)
	paymentService := services.NewPaymentService(invoiceRepo, paymentRepo, exchangeRateRepo)
	invoiceGenService := services.NewInvoiceGenerationService(
		enrollmentRepo, feeStructureRepo, feeInstallmentRepo, invoiceRepo, sectionRepo,
	)

	return &Application{
		Config:            cfg,
		DB:                db,
		SectionService:    sectionService,
		StudentService:    studentService,
		EnrollmentService: enrollmentService,
		PaymentService:    paymentService,
		InvoiceGenService: invoiceGenService,
		InvoiceRepo:       invoiceRepo,
		PaymentRepo:       paymentRepo,
	}
}
