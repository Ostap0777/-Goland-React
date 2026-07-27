package makes

import (
	"backend/internal/httputil"
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	store *Store
}

func NewHandler(s *Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/makes", h.list)
	mux.HandleFunc("POST /api/makes", h.create)
	mux.HandleFunc("GET /api/makes/{id}", h.get)
	mux.HandleFunc("DELETE /api/makes/{id}", h.delete)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	makes, err := h.store.List()
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to list makes")
		return
	}
	httputil.WriteJSON(w, http.StatusOK, makes)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var in Input
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&in); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return
	}
	if err := in.Validate(); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	m, err := h.store.Create(in)
	if errors.Is(err, ErrDuplicate) {
		httputil.WriteError(w, http.StatusConflict, "make already exists")
		return
	}
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to create make")
		return
	}
	httputil.WriteJSON(w, http.StatusCreated, m)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, ok := httputil.ParseID(w, r)
	if !ok {
		return
	}
	m, err := h.store.Get(id)
	if errors.Is(err, ErrNotFound) {
		httputil.WriteError(w, http.StatusNotFound, "make not found")
		return
	}
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to get make")
		return
	}
	httputil.WriteJSON(w, http.StatusOK, m)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := httputil.ParseID(w, r)
	if !ok {
		return
	}
	if err := h.store.Delete(id); err != nil {
		if errors.Is(err, ErrNotFound) {
			httputil.WriteError(w, http.StatusNotFound, "make not found")
			return
		}
		httputil.WriteError(w, http.StatusInternalServerError, "failed to delete make")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
