package repositories

import (
	"github.com/Kenini1805/go-rest-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	IsDuplicateEmail(email string) (bool, error)
	InsertUser(user models.User) (models.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) IsDuplicateEmail(email string) (bool, error) {
	var user models.User
	err := db.connection.Where("email = ?", email).First(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *userConnection) InsertUser(user models.User) (models.User, error) {
	err := db.connection.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
