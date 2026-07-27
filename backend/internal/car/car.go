package car

import (
	"errors"
	"strings"
	"time"
)

type Car struct {
	ID    int64 `json:"id"`
	Make  string `json:"make"`//e.g "Toyota"
	Model string `json:"model"`//e.g "Corolla"
	Year int `json:"yaer"`//2015
	Price int64 `json:"price"`//20.000$
	Mileage  int64 `json:"mileage"`//odometr reading, in km
	Description string `json:"description"`//nice car, cool
	SellerName string `json:"sellecrName"`//Who published this car
	CreatedAt time.Time `json:"createdAt"`//20.05.2025
}


type Input struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year int `json:"year"`
	Price int64 `json:"price"`
	Mileage  int64 `json:"mileage"`
	Description string `json:"description"`
	SellerName string `json:"sellerName"`
}


//guard
const firstCarYear = 1886


func (in *Input) Validate() error  {
	in.Make = strings.TrimSpace(in.Make)
	in.Model = strings.TrimSpace(in.Model)
	in.Description = strings.TrimSpace(in.Description)
	in.SellerName = strings.TrimSpace(in.SellerName)


	switch{
	case in.Make == "":
		return errors.New("make is required")
	case in.Model == "":
		return errors.New("model is required")
	case in.SellerName == "":
	return  errors.New("setterName is required")
	
case in.Year < firstCarYear || in.Year > time.Now().Year()+1:
	return errors.New("year is out of range")
case in.Price < 0:
	return errors.New("price must be positive")
case in.Mileage < 0:
	return errors.New("mileage cannot be negative")
}
return nil
}