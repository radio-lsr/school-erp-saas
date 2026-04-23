package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/radio-lsr/school-erp-saas/backend/internal/app"
)

type AuthHandler struct {
    app *app.Application
}

func NewAuthHandler(app *app.Application) *AuthHandler {
    return &AuthHandler{app: app}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": "fake-jwt-token",
        "message": "Login endpoint",
    })
}