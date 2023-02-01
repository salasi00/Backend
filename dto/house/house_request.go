package housedto

import "housy/models"

type CreateHouseRequest struct {
	ID        int         `json:"id"`
	Name      string      `json:"name" form:"name" validate:"required"`
	CityID    int         `json:"cityid" form:"cityid" validate:"required"`
	City      models.City `json:"-"`
	Address   string      `json:"address" form:"address" validate:"required"`
	Price     int         `json:"price" form:"price" validate:"required"`
	TypeRent  string      `json:"typerent" form:"typerent" validate:"required"`
	Amenities string      `json:"amenities" form:"amenities" validate:"required"`
	BedRoom   int         `json:"bedRoom" form:"bedRoom" validate:"required"`
	Bathroom  int         `json:"bathroom" form:"bathroom" validate:"required"`
	Image     string      `json:"image" form:"image"`
}

type UpdateHouseRequest struct {
	Name      string `json:"name" form:"name"`
	CityID    int    `json:"cityid" form:"cityid"`
	Address   string `json:"address" form:"address"`
	Price     string `json:"price" form:"price"`
	TypeRent  string `json:"typerent" form:"typerent"`
	Amenities string `json:"aminities" form:"amenities"`
	BedRoom   string `json:"bedRoom" form:"bedRoom"`
	Bathroom  string `json:"bathroom" form:"bathroom"`
}
