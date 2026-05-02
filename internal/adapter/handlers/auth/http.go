package auth

import (
    "encoding/json"
    "net/http"
    userport "github.com/bishal-dhakal/project-management/internal/core/port/user"
)

type Handler struct {
    svc userport.AuthService
}

func NewHandler(svc userport.AuthService) *Handler {
    return &Handler{svc: svc}
}

// request struct — maps incoming JSON to Go struct
type authRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
    // 1. decode JSON body → get email and password
    var req authRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    // 2. now email and password exist — call service
    data, err := h.svc.Register(r.Context(), req.Email, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 3. return response
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "message": data,
    })
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    var req authRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    token, err := h.svc.Login(r.Context(), req.Email, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}