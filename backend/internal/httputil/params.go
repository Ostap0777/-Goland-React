package httputil

import (
	"backend/internal/httputil"
	"net/http"
	"strconv"
)

func ParseID(w http.ResponseWriter, r *http.Request)(int64,bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10,64)
	if err != nil || id <=0 {
		httputil.WriteError(w, http.StatusBadRequest, "invalid car id")
		return 0, false
	}
	return id, true
}