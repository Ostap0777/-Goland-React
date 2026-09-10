package car

import "time"

type CreateCarRequest struct {
    Make        string            `json:"make"`
    Model       string            `json:"model"`
    Year        int               `json:"year"`
    Price       int64             `json:"price"`
    Mileage     int64             `json:"mileage"`
    Description string            `json:"description"`
    SellerName  string            `json:"sellerName"`
    Images      []CreateCarImage  `json:"images"`
}

type CarResponse struct {
    ID          int64              `json:"id"`
    Make        string             `json:"make"`
    Model       string             `json:"model"`
    Year        int                `json:"year"`
    Price       int64              `json:"price"`
    Mileage     int64              `json:"mileage"`
    Description string             `json:"description"`
    SellerName  string             `json:"sellerName"`
    Images      []CarImageResponse `json:"images,omitempty"`
    CreatedAt   string             `json:"createdAt"`
}

type CreateCarImage struct {
	URL    string `json:"url"`
	IsMain bool   `json:"isMain"`
	Order  int    `json:"order"`
}

type CarImageResponse struct {
	ID        int64  `json:"id"`
	URL       string `json:"url"`
	IsMain    bool   `json:"isMain"`
	Order     int    `json:"order"`
	CreatedAt string `json:"createdAt"`
}

func toResponse(c *Car) CarResponse {
	images := make([]CarImageResponse, 0, len(c.Images))
	for _, img := range c.Images {
		images = append(images,CarImageResponse{
			ID: img.ID,
			URL: img.URL,
			IsMain: img.IsMain,
			Order: img.Order,
			CreatedAt: img.CreatedAt.UTC().Format(time.RFC3339),
		})
	}
	return CarResponse{
		ID:          c.ID,
		Make:        c.Make,
		Model:       c.Model,
		Year:        c.Year,
		Price:       c.Price,
		Mileage:     c.Mileage,
		Description: c.Description,
		Images: images,
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
