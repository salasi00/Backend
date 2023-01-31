package authdto

import "housy/models"

type RegisterRequest struct {
	FullName string                 `json:"fullName" gorm:"type varchar(255)" validate:"required"`
	Username string                 `json:"username" gorm:"type varchar(255)" validate:"required"`
	Email    string                 `json:"email" gorm:"type varchar(255)" validate:"required"`
	Password string                 `json:"password" gorm:"type varchar(255)" validate:"required"`
	ListAsID int                    `json:"listasid"`
	Profile  models.ProfileResponse `json:"profile"`
}

type LoginRequest struct {
	Username string `json:"username" gorm:"type varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type varchar(255)" validate:"required"`
}
