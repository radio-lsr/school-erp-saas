package middleware

import (
    "context"
    "net/http"
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/config"
)

type ContextKey string

const (
    TenantIDKey ContextKey = "tenantID"
    UserIDKey   ContextKey = "userID"
)

func AuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Pour l'instant, bypass complet
            ctx := context.WithValue(r.Context(), TenantIDKey, uuid.MustParse("00000000-0000-0000-0000-000000000001"))
            ctx = context.WithValue(ctx, UserIDKey, uuid.MustParse("00000000-0000-0000-0000-000000000001"))
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// Si vous activez JWT, décommentez et implémentez :
// func parseToken(...)