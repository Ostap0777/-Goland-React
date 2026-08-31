package car

import (
	"strings"
	"time"
)

func (in *CreateCarRequest) Validate() error {
	in.Make = strings.TrimSpace(in.Make)
	in.Model = strings.TrimSpace(in.Model)
	in.Description = strings.TrimSpace(in.Description)
	in.SellerName = strings.TrimSpace(in.SellerName)

	switch {
	case in.Make == "":
		return &InputError{Message: "make is required"}
	case in.Model == "":
		return &InputError{Message: "model is required"}
	case in.SellerName == "":
		return &InputError{Message: "sellerName is required"}
	case in.Year < firstCarYear || in.Year > time.Now().Year()+1:
		return &InputError{Message: "year is out of range"}
	case in.Price < 0:
		return &InputError{Message: "price cannot be negative"}
	case in.Mileage < 0:
		return &InputError{Message: "mileage cannot be negative"}
	}
	return nil
}
