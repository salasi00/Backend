package authdto

type RegisterRequest struct {
	Username string `json:"username" gorm:"type varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type varchar(255)" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" gorm:"type varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type varchar(255)" validate:"required"`
}
