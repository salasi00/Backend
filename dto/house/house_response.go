package housedto

import "housy/models"

type HouseResponse struct {
	ID        int         `json:"id"`
	Name      string      `json:"name" form:"name" validate:"required"`
	CityID    int         `json:"-" validate:"required"`
	City      models.City `json:"city"`
	Address   string      `json:"address" form:"address" validate:"required"`
	Price     int         `json:"price" form:"price" validate:"required"`
	TypeRent  string      `json:"typerent" form:"typerent" validate:"required"`
	Amenities string      `json:"amenities" form:"amenities" validate:"required"`
	BedRoom   int         `json:"bedRoom" form:"bedRoom" validate:"required"`
	Bathroom  int         `json:"bathroom" form:"bathroom" validate:"required"`
}
