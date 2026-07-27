package car

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
	mux.HandleFunc("GET /api/cars", h.list)
	mux.HandleFunc("POST /api/cars", h.create)
	mux.HandleFunc("PUT /api/cars/{id}", h.update)
	mux.HandleFunc("GET /api/cars/{id}", h.get)
	mux.HandleFunc("DELETE /api/cars/{id}", h.delete)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	cars, err := h.store.List()
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to list cars")
		return
	}
	httputil.WriteJSON(w, http.StatusOK, cars)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, ok := httputil.ParseID(w,r)
	if !ok {
		return
	}
	m, err := h.store.Get(id)
	if errors.Is(err, ErrNotFound) {
		httputil.WriteError(w, http.StatusNotFound, "car not found")
		return
	}
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to list cars")
		return
	}
	httputil.WriteJSON(w, http.StatusOK, m)
}


func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var in Input
	dec = json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&in); err != nil {
		httputil.WriteError(w, http.StatusBadRequest,"invalid JSON body: "+err.Error())
		return
	}
	if err := in.Validate(); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	m, err := h.store.Create(in)
	if error.Is(err, ErrDuplicate) {
		httputil.WriteError(w, http.StatusConflict, "car already exists")
		return
	}
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to create car")
		return
	}
	httputil.WriteJSON(w, http.StatusCreated, m)
}


func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, ok := httputil.ParseID(w,r)
	if !ok {
		return
	}
	dec = json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&in); err != nil {
		httputil.WriteError(w, http.StatusBadRequest,"invalid JSON body: "+err.Error())
		return
	}
	if err := in.Validate(); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	car, err := h.store.Update(id, in)
		if error.Is(err, ErrNotFound) {
		httputil.WriteError(w, http.StatusNotFound, "car not found")
		return
	}
	if error.Is(err, ErrDuplicate) {
		httputil.WriteError(w, http.StatusConflict, "car already exists")
		return
	}
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to create car")
		return
	}
	httputil.WriteJSON(w, http.StatusOK, car)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Ruquest) {
	id, ok := httputil.ParseID(w,r)
	if !ok {
		return
	}
	if err := h.store.Delete(id); err != nil {
		if errors.Is(err, ErrNotFound) {
			httputil.WriteError(w, http.StatusNotFound, "cars not found")
			return
		}
		httputil.WriteError(w, http.StatusInternalServerError, "failed to delete car")
	}
	w.WriteHeader(http.StatusNoContent)
}
