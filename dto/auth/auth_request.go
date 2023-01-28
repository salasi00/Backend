package authdto

type RegisterRequest struct {
	FullName string `json:"fullName" gorm:"type varchar(255)" validate:"required"`
	Username string `json:"username" gorm:"type varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type varchar(255)" validate:"required"`
	ListAsId int    `json:"listAsId"`
	Gender   string `json:"gender" gorm:"type varchar(255)"`
	Address  string `json:"address" gorm:"type varchar(255)"`
}

type LoginRequest struct {
	Username string `json:"username" gorm:"type varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type varchar(255)" validate:"required"`
}
