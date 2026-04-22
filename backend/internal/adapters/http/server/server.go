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
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: true,
    }))

    // Route publique
    r.Post("/api/auth/login", handlers.NewAuthHandler(application).Login)

    // Routes protégées (à activer plus tard)
    r.Group(func(r chi.Router) {
        r.Use(middleware.AuthMiddleware(cfg))
        // r.Mount("/api/students", handlers.NewStudentHandler(application.StudentService).Routes())
        // r.Mount("/api/sections", handlers.NewSectionHandler(application.SectionService).Routes())
    })

    return &http.Server{
        Addr:    ":" + cfg.Port,
        Handler: r,
    }
}