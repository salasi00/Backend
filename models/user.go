package models

import "time"

type User struct {
	ID        int
	FullName  string    `json:"fullName" gorm:"type: varchar(255)"`
	Username  string    `json:"username" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	ListAsID  int       `json:"list_as_id"`
	ListAs    string    `json:"list_as" gorm:"tyle: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
