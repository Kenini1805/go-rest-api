package controllers

import (
	"errors"
	"net/http"

	"github.com/Kenini1805/go-rest-api/internal/models"
	"github.com/Kenini1805/go-rest-api/internal/resources"
	"github.com/Kenini1805/go-rest-api/internal/services"
	"github.com/Kenini1805/go-rest-api/pkg/converter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController interface {
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Register godoc
// @Summary Register user
// @Description Register user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 500 {object} httperrors.RestError
// @Router /auth/register [post]
func (c *authController) Register(ctx *gin.Context) {
	var registerRequest models.RegisterUserRequest
	errRequest := ctx.ShouldBind(&registerRequest)
	if errRequest != nil {
		response := converter.BuildErrorResponse("Failed to process the request", errRequest.Error(), converter.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userExists, err := c.authService.IsDuplicateEmail(registerRequest.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response := converter.BuildErrorResponse("Failed to process the request", err.Error(), converter.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}

	if userExists {
		response := converter.BuildErrorResponse("Failed to process the request", "Duplicate email", converter.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)

		return
	}
	user, err := c.authService.CreateUser(registerRequest)
	if err != nil {
		response := converter.BuildErrorResponse("Failed to process the request", err.Error(), converter.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)

		return
	}
	// TO DO Send Email
	ctx.JSON(http.StatusOK, resources.NewUserResponse(user))
}
