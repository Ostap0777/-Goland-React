package car

import (
	"backend/internal/httputil"
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

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/cars", h.list)
	mux.HandleFunc("POST /api/cars", h.create)
	mux.HandleFunc("GET /api/cars/{id}", h.get)
	mux.HandleFunc("PUT /api/cars/{id}", h.update)
	mux.HandleFunc("DELETE /api/cars/{id}", h.delete)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	filter, err := ParseFilter(r)
	if err != nil {
		httputil.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cars, err := h.service.ListCars(r.Context(), filter)
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to list cars")
		return
	}
	httputil.WriteJSON(w, http.StatusOK, toResponses(cars))
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, ok := httputil.ParseID(w, r)
	if !ok {
		return
	}

	c, err := h.service.GetCar(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		httputil.WriteError(w, http.StatusNotFound, "car not found")
		return
	}
	if err != nil {
		httputil.WriteError(w, http.StatusInternalServerError, "failed to get car")
		return
	}
	httputil.WriteJSON(w, http.StatusOK, toResponse(c))
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	in, ok := decodeCarRequest(w, r)
	if !ok {
		return
	}

	c, err := h.service.CreateCar(r.Context(), in)
	if !writeCarError(w, err, "failed to create car") {
		return
	}
	httputil.WriteJSON(w, http.StatusCreated, toResponse(c))
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, ok := httputil.ParseID(w, r)
	if !ok {
		return
	}

	in, ok := decodeCarRequest(w, r)
	if !ok {
		return
	}

	c, err := h.service.UpdateCar(r.Context(), id, in)
	if !writeCarError(w, err, "failed to update car") {
		return
	}
	httputil.WriteJSON(w, http.StatusOK, toResponse(c))
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := httputil.ParseID(w, r)
	if !ok {
		return
	}
	if err := h.service.DeleteCar(r.Context(), id); err != nil {
		if errors.Is(err, ErrNotFound) {
			httputil.WriteError(w, http.StatusNotFound, "car not found")
			return
		}
		httputil.WriteError(w, http.StatusInternalServerError, "failed to delete car")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodeCarRequest(w http.ResponseWriter, r *http.Request) (CreateCarRequest, bool) {
	var in CreateCarRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&in); err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return CreateCarRequest{}, false
	}
	return in, true
}

func writeCarError(w http.ResponseWriter, err error, fallback string) bool {
	if err == nil {
		return true
	}

	var inputErr *InputError
	switch {
	case errors.As(err, &inputErr):
		httputil.WriteError(w, http.StatusBadRequest, inputErr.Error())
	case errors.Is(err, ErrUnknownMake):
		httputil.WriteError(w, http.StatusBadRequest, "unknown make")
	case errors.Is(err, ErrNotFound):
		httputil.WriteError(w, http.StatusNotFound, "car not found")
	default:
		httputil.WriteError(w, http.StatusInternalServerError, fallback)
	}
	return false
}
