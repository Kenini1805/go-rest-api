package services

import (
	"fmt"
	"log"

	"github.com/Kenini1805/go-rest-api/internal/models"
	"github.com/Kenini1805/go-rest-api/internal/repositories"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	IsDuplicateEmail(email string) (bool, error)
	CreateUser(userRegister models.RegisterUserRequest) (models.User, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRep repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) IsDuplicateEmail(email string) (bool, error) {
	res, err := service.userRepository.IsDuplicateEmail(email)
	if err != nil {
		return false, fmt.Errorf("failed to check duplicate email: %w", err)
	}

	return res, err
}

func (service *authService) CreateUser(userRegister models.RegisterUserRequest) (models.User, error) {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&userRegister))
	if err != nil {
		return models.User{}, fmt.Errorf("fail to register user: %w", err)
	}
	user.Password = hashAndSalt([]byte(user.Password))
	res, err := service.userRepository.InsertUser(user)
	if err != nil {
		return models.User{}, fmt.Errorf("fail to register user: %w", err)
	}

	return res, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed hash password")
	}

	return string(hash)
}
