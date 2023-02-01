package models

type House struct {
	ID        int    `json:"id"`
	Name      string `json:"name" form:"name" gorm:"type: varchar(255)"`
	CityID    int    `json:"-"`
	City      City   `json:"city"`
	Address   string `json:"address" gorm:"type: varchar(255)" form:"address"`
	Price     int    `json:"price" gorm:"type: int" form:"price:"`
	TypeRent  string `json:"typerent" gorm:"type: varchar(255)" form:"typeRent"`
	Amenities string `json:"amenities" gorm:"type: varchar(255)" form:"aminities"`
	BedRoom   int    `json:"bedRoom" gorm:"type: int" form:"bedRoom"`
	Bathroom  int    `json:"bathroom" gorm:"type: int" form:"bathroom"`
	Image     string `json:"image" form:"image"`
	UserID    int    `json:"user_id"`
	User      UsersProfileResponse
}

func (House) TableName() string {
	return "houses"
}
