package models

// type Product struct {
// 	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
// 	Name       string               `json:"name" form:"name" gorm:"type: varchar(255)"`
// 	Desc       string               `json:"desc" form:"desc" gorm:"type: text"`
// 	Price      int                  `json:"price" form:"price" gorm:"type: int"`
// 	Image      string               `json:"image" form:"image" gorm:"type: varchar(255)"`
// 	Qty        int                  `json:"qty" form:"qty" gorm:"type: int"`
// 	UserID     int                  `json:"user_id" form:"user_id" gorm:"type int"`
// 	User       UsersProfileResponse `json:"user"`
// 	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
// 	CategoryID []int                `json:"category_id" form :"category_id" gorm:"-"`
// }

// type ProductResponse struct {
// 	ID         int                  `json:"id"`
// 	Name       string               `json:"name"`
// 	Desc       string               `json:"desc"`
// 	Price      int                  `json:"price"`
// 	Image      string               `json:"image"`
// 	Qty        int                  `json:"qty"`
// 	UserID     int                  `json:"-"`
// 	User       UsersProfileResponse `json:"user"`
// 	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
// 	CategoryID []int                `json:"category_id" form:"category_id" gorm:"-"`
// }

// type ProductUserResponse struct {
// 	ID     int    `json:"id"`
// 	Name   int    `json:"name"`
// 	Desc   string `json:"desc"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	Qty    int    `json:"qty"`
// 	UserID int    `json:"-"`
// }

// func (ProductResponse) TableName() string {
// 	return "products"
// }

// func (ProductUserResponse) TableName() string {
// 	return "products"
// }
