package app

import (
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/radio-lsr/school-erp-saas/backend/internal/config"
    "github.com/radio-lsr/school-erp-saas/backend/internal/core/services"
    "github.com/radio-lsr/school-erp-saas/backend/internal/adapters/repository"
)

type Application struct {
    Config          *config.Config
    DB              *pgxpool.Pool
    SectionService  *services.SectionService
    StudentService  *services.StudentService
}

func NewApplication(db *pgxpool.Pool, cfg *config.Config) *Application {
    // Repositories
    sectionRepo := repository.NewPostgresSectionRepository(db)
    studentRepo := repository.NewPostgresStudentRepository(db)

    // Services
    sectionService := services.NewSectionService(sectionRepo)
    studentService := services.NewStudentService(studentRepo)

    return &Application{
        Config:         cfg,
        DB:             db,
        SectionService: sectionService,
        StudentService: studentService,
    }
}