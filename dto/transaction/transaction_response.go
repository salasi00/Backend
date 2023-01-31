package transactiondto

import (
	"housy/models"
)

type TransactionResponse struct {
	ID         int          `json:"id"`
	CheckIn    string       `json:"checkin"`
	CheckOut   string       `json:"checkout"`
	HouseID    int          `json:"-"`
	House      models.House `json:"house"`
	Total      int          `json:"total"`
	Status     string       `json:"status"`
	Attachment string       `json:"attachment"`
}
