package authdto

type LoginResponse struct {
	Username string `json:"username" gorm:"type varchar(255)"`
	Token    string `json:"token" gorm:"type varchar(255)"`
}
