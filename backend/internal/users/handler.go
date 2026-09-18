package users

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/user/{id}", h.getByID)
	// mux.HandleFunc("GET /api/user/{id}", h.getByEmail)
	// mux.HandleFunc("PUT /api/user/{email}", h.update)
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	idSrt := r.PathValue("id")
	id, err := strconv.ParseInt(idSrt, 10,64)
	if err != nil {
		http.Error(w, "Invelid user ID", http.StatusBadRequest)
	}
	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if err == ErrUserNotFound {
			http.Error(w,"User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ToUserResponse(user)); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
