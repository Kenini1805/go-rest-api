package controllers

import (
	"errors"
	"net/http"

	"github.com/Kenini1805/go-rest-api/internal/models"
	"github.com/Kenini1805/go-rest-api/internal/resources"
	"github.com/Kenini1805/go-rest-api/internal/services"
	"github.com/Kenini1805/go-rest-api/pkg/converter"
	httperrors "github.com/Kenini1805/go-rest-api/pkg/http_errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
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
	if err := BindRequest(ctx, &registerRequest); err != nil {
		return
	}

	userExists, err := c.authService.IsDuplicateEmail(registerRequest.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response := converter.BuildErrorResponse(httperrors.ErrBadRequestMessage, err.Error(), converter.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}

	if userExists {
		response := converter.BuildErrorResponse(httperrors.ErrBadRequestMessage, "Duplicate email", converter.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)

		return
	}
	user, err := c.authService.CreateUser(registerRequest)
	if err != nil {
		response := converter.BuildErrorResponse(httperrors.ErrBadRequestMessage, err.Error(), converter.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)

		return
	}
	// TO DO Send Email
	ctx.JSON(http.StatusOK, resources.NewUserResponse(user))
}

// Login godoc
// @Summary Login
// @Description Login with JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 string access_token
// @Failure 500 {object} httperrors.RestError
// @Router /auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	if err := BindRequest(ctx, &loginRequest); err != nil {
		return
	}

	authUser, err := c.authService.VerifyCredential(loginRequest)
	if err != nil {
		response := converter.BuildErrorResponse(httperrors.ErrBadRequestMessage, err.Error(), converter.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	if authUser.Email != "" {
		generateToken := c.jwtService.GenerateToken(authUser.ID.String())
		ctx.JSON(http.StatusOK, resources.NewLoginResponse(generateToken))

		return
	}
	response := converter.BuildErrorResponse(
		httperrors.ErrBadRequestMessage,
		httperrors.ErrUnauthorized.Error(), converter.EmptyObj{},
	)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
