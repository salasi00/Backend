package transactiondto

import (
	"housy/models"
)

type TransactionRequest struct {
	CheckIn    string       `json:"checkin" form:"checkin" validate:"required"`
	CheckOut   string       `json:"checkout" form:"checkout" validate:"required"`
	HouseID    int          `json:"houseid" form:"houseid" validate:"required"`
	House      models.House `json:"house"`
	UserID     int          `json:"userid" form:"userid" validate:"required"`
	User       models.User  `json:"user"`
	Total      string       `json:"total" form:"total" validate:"required"`
	Status     string       `json:"status" form:"status" validate:"required"`
	Attachment string       `json:"attachment" form:"attachment" validate:"required"`
}
