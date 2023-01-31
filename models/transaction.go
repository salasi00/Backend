package models

type Transaction struct {
	ID         int                  `json:"id"`
	CheckIn    string               `json:"checkin"`
	CheckOut   string               `json:"checkout"`
	HouseID    int                  `json:"houseid"`
	House      House                `json:"house"`
	UserID     int                  `json:"userid"`
	User       UsersProfileResponse `json:"user"`
	Total      string               `json:"email" gorm:"type: varchar(255)"`
	Status     string               `json:"status" gorm:"type: varchar(255)"`
	Attachment string               `json:"attachment" gorm:"type: varchar(255)"`
}
