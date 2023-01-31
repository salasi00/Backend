package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(username string) (models.User, error)
	CreateProfile(user models.User) error
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) CreateProfile(user models.User) error {
	profile := models.Profile{
		Gender:  user.Profile.Gender,
		Phone:   user.Profile.Phone,
		Address: user.Profile.Address,
		UserID:  user.ID,
	}
	err := r.db.Create(&profile).Error

	return err
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error

	return user, err
}
