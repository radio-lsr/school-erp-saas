package middleware

import (
    "context"
    "net/http"
    "github.com/google/uuid"
)

type ContextKey string

const (
    TenantIDKey ContextKey = "tenantID"
    UserIDKey   ContextKey = "userID"
)

func AuthMiddleware(cfg interface{}) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Pour l'instant, bypass complet (fake tenant/user)
            ctx := context.WithValue(r.Context(), TenantIDKey, uuid.MustParse("00000000-0000-0000-0000-000000000001"))
            ctx = context.WithValue(ctx, UserIDKey, uuid.MustParse("00000000-0000-0000-0000-000000000001"))
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}