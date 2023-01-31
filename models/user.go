package models

import "time"

type User struct {
	ID        int
	FullName  string          `json:"fullName" gorm:"type: varchar(255)"`
	Username  string          `json:"username" gorm:"type: varchar(255)"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"password" gorm:"type: varchar(255)"`
	Profile   ProfileResponse `json:"profile" gorm:"constraint:OnUpdate:CASCADE"`
	ListAsID  int             `json:"-"`
	ListAs    ListAs          `json:"list_as"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
