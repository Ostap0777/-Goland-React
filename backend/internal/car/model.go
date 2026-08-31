package car

import "time"

const firstCarYear = 1886

type Car struct {
	ID          int64
	Make        string
	Model       string
	Year        int
	Price       int64
	Mileage     int64
	Description string
	SellerName  string
	Images      []CarImage `json:"images,omitempty"`
	CreatedAt   time.Time
}

type CarImage struct {
	ID        int64     `json:"id"`
	CarID     int64     `json:"car_id"`
	URL       string    `json:"url"`
	IsMain    bool      `json:"is_main"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
}
