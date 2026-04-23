package app

import (
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/config"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
    "github.com/radio-lsr/school-erp-saas/backend/internal/adapters/repository"
)

type Application struct {
    Config             *config.Config
    DB                 *pgxpool.Pool
    SectionService     *services.SectionService
    StudentService     *services.StudentService
    EnrollmentService  *services.EnrollmentService
    PaymentService     *services.PaymentService
    InvoiceGenService  *services.InvoiceGenerationService
}

func NewApplication(db *pgxpool.Pool, cfg *config.Config) *Application {
    sectionRepo := repository.NewPostgresSectionRepository(db)
    studentRepo := repository.NewPostgresStudentRepository(db)
    enrollmentRepo := repository.NewPostgresEnrollmentRepository(db)
    feeStructRepo := repository.NewPostgresFeeStructureRepository(db)
    feeInstallmentRepo := repository.NewPostgresFeeInstallmentRepository(db)
    invoiceRepo := repository.NewPostgresInvoiceRepository(db)
    paymentRepo := repository.NewPostgresPaymentRepository(db)
    exchangeRateRepo := repository.NewPostgresExchangeRateRepository(db)

    sectionService := services.NewSectionService(sectionRepo)
    studentService := services.NewStudentService(studentRepo)
    enrollmentService := services.NewEnrollmentService(enrollmentRepo, sectionRepo, studentRepo)
    paymentService := services.NewPaymentService(invoiceRepo, paymentRepo, exchangeRateRepo)
    invoiceGenService := services.NewInvoiceGenerationService() // no dependencies yet

    return &Application{
        Config:            cfg,
        DB:                db,
        SectionService:    sectionService,
        StudentService:    studentService,
        EnrollmentService: enrollmentService,
        PaymentService:    paymentService,
        InvoiceGenService: invoiceGenService,
    }
}