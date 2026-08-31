package car

import "time"

type CreateCarRequest struct {
	Make        string `json:"make"`
	Model       string `json:"model"`
	Year        int    `json:"year"`
	Price       int64  `json:"price"`
	Mileage     int64  `json:"mileage"`
	Description string `json:"description"`
	SellerName  string `json:"sellerName"`
}

type CarResponse struct {
	ID          int64  `json:"id"`
	Make        string `json:"make"`
	Model       string `json:"model"`
	Year        int    `json:"year"`
	Price       int64  `json:"price"`
	Mileage     int64  `json:"mileage"`
	Description string `json:"description"`
	SellerName  string `json:"sellerName"`
	CreatedAt   string `json:"createdAt"`
}

func toResponse(c *Car) CarResponse {
	return CarResponse{
		ID:          c.ID,
		Make:        c.Make,
		Model:       c.Model,
		Year:        c.Year,
		Price:       c.Price,
		Mileage:     c.Mileage,
		Description: c.Description,
		SellerName:  c.SellerName,
		CreatedAt:   c.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func toResponses(cars []Car) []CarResponse {
	out := make([]CarResponse, 0, len(cars))
	for i := range cars {
		out = append(out, toResponse(&cars[i]))
	}
	return out
}
