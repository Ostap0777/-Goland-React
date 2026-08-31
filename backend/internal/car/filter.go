package car

import (
	"fmt"
	"net/http"
	"strings"

	"backend/internal/httputil"
)

type Filter struct {
	Make     string
	Model    string
	MinYear  int
	MaxYear  int
	MinPrice int64
	MaxPrice int64
	Limit    int
	Offset   int
}

func ParseFilter(r *http.Request) (Filter, error) {
	q := r.URL.Query()
	f := Filter{
		Make:  strings.TrimSpace(q.Get("make")),
		Model: strings.TrimSpace(q.Get("model")),
	}

	if v, ok, err := httputil.QueryInt(r, "minYear"); err != nil {
		return Filter{}, fmt.Errorf("invalid minYear")
	} else if ok {
		f.MinYear = v
	}

	if v, ok, err := httputil.QueryInt(r, "maxYear"); err != nil {
		return Filter{}, fmt.Errorf("invalid maxYear")
	} else if ok {
		f.MaxYear = v
	}

	if v, ok, err := httputil.QueryInt64(r, "minPrice"); err != nil {
		return Filter{}, fmt.Errorf("invalid minPrice")
	} else if ok {
		f.MinPrice = v
	}

	if v, ok, err := httputil.QueryInt64(r, "maxPrice"); err != nil {
		return Filter{}, fmt.Errorf("invalid maxPrice")
	} else if ok {
		f.MaxPrice = v
	}

	if v, ok, err := httputil.QueryInt(r, "limit"); err != nil {
		return Filter{}, fmt.Errorf("invalid limit")
	} else if ok {
		f.Limit = v
	}

	if v, ok, err := httputil.QueryInt(r, "offset"); err != nil {
		return Filter{}, fmt.Errorf("invalid offset")
	} else if ok {
		f.Offset = v
	}

	return f, nil
}
