package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/Kenini1805/go-rest-api/internal/models"
	"github.com/Kenini1805/go-rest-api/internal/repositories"
	"github.com/gofrs/uuid"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	IsDuplicateEmail(email string) (bool, error)
	CreateUser(userRegister models.RegisterUserRequest) (models.User, error)
	VerifyCredential(credentials models.LoginRequest) (models.User, error)
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
	user.ID = uuid.Must(uuid.NewV4())
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

func (service *authService) VerifyCredential(credentials models.LoginRequest) (models.User, error) {
	password := credentials.Password
	res, err := service.userRepository.VerifyCredential(credentials)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, nil
		}
		return models.User{}, fmt.Errorf("fail to verify credential user: %w", err)
	}
	comparedPassword := comparePassword(res.Password, []byte(password))

	if !comparedPassword {
		return models.User{}, nil
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

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
