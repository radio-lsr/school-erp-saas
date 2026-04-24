package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "github.com/radio-lsr/school-erp-saas/backend/internal/app"
)

type AuthHandler struct {
    app *app.Application
}

func NewAuthHandler(app *app.Application) *AuthHandler {
    return &AuthHandler{app: app}
}

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req loginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    // Récupérer l'utilisateur par email
    user, err := h.app.UserRepo.GetByEmail(r.Context(), req.Email)
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    if user == nil {
        http.Error(w, "invalid credentials", http.StatusUnauthorized)
        return
    }

    // Vérifier le mot de passe (stocké en clair pour démo, à remplacer par bcrypt plus tard)
    if user.PasswordHash != req.Password {
        http.Error(w, "invalid credentials", http.StatusUnauthorized)
        return
    }

    // Créer le JWT
    claims := jwt.MapClaims{
        "user_id":  user.ID.String(),
        "tenant_id": user.TenantID.String(),
        "role":     user.Role,
        "exp":      time.Now().Add(72 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(h.app.Config.JWTSecret))
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    // Réponse
    resp := map[string]interface{}{
        "token": tokenString,
        "user": map[string]string{
            "id":        user.ID.String(),
            "email":     user.Email,
            "fullName":  user.FullName,
            "role":      user.Role,
            "tenantId":  user.TenantID.String(),
        },
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}