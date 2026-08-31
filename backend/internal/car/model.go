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
	CreatedAt   time.Time
}
