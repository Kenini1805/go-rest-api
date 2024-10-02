package resources

import (
	"github.com/Kenini1805/go-rest-api/internal/models"
)

type UserResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		UserName: user.UserName,
		Email:    user.Email,
		Role:     *user.Role,
	}
}
