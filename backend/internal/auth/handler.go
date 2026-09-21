package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
   mux.HandleFunc("POST /api/auth/register", h.register)
	mux.HandleFunc("GET /api/user/{id}", h.getByEmail)
	// mux.HandleFunc("PUT /api/user/{email}", h.update)
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
var req RegisterRequest

	// 1. Декодування тіла запиту
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// 2. Виклики бізнес-логіки
	userResp, err := h.service.Register(r.Context(), req)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			http.Error(w, "Email already in use", http.StatusConflict)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 3. Успішна відповідь (201 Created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResp)
}

func (h *Handler) getByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
       http.Error(w, "Email parameter is required", http.StatusBadRequest)
       return
	}

	userResp, err := h.service.GetByEmail(r.Context(), email)
	if err != nil {
		if errors.Is(err, ErrEmailNotFound) {
			http.Error(w, "Email not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResp)
}

