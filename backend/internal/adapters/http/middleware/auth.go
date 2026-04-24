package middleware

import (
    "context"
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "github.com/radio-lsr/school-erp-saas/backend/internal/config"
)

type ContextKey string

const (
    TenantIDKey ContextKey = "tenantID"
    UserIDKey   ContextKey = "userID"
    RoleKey     ContextKey = "role"
)

func AuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "missing authorization header", http.StatusUnauthorized)
                return
            }

            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                return []byte(cfg.JWTSecret), nil
            })
            if err != nil || !token.Valid {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }

            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok {
                http.Error(w, "invalid claims", http.StatusUnauthorized)
                return
            }

            tenantID, err := uuid.Parse(claims["tenant_id"].(string))
            if err != nil {
                http.Error(w, "invalid tenant in token", http.StatusUnauthorized)
                return
            }
            userID, err := uuid.Parse(claims["user_id"].(string))
            if err != nil {
                http.Error(w, "invalid user in token", http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), TenantIDKey, tenantID)
            ctx = context.WithValue(ctx, UserIDKey, userID)
            if role, ok := claims["role"].(string); ok {
                ctx = context.WithValue(ctx, RoleKey, role)
            }
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}