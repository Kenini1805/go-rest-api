package server

import (
	"github.com/Kenini1805/go-rest-api/internal/controllers"
	"github.com/Kenini1805/go-rest-api/internal/repositories"
	"github.com/Kenini1805/go-rest-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	userRepository repositories.UserRepository
	authService    services.AuthService
	jwtService     services.JWTService
	authController controllers.AuthController
}

// NewHandlers initializes the handlers with the necessary dependencies.
func NewHandlers(db *gorm.DB) *Handlers {
	// Initialize repositories, services, and controllers with db.
	userRepository := repositories.NewUserRepository(db)
	jwtService := services.NewJWTService()
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService, jwtService)

	// Return the Handlers struct with all the initialized dependencies.
	return &Handlers{
		userRepository: userRepository,
		authService:    authService,
		jwtService:     jwtService,
		authController: authController,
	}
}

// MapRoutes defines the routes for the server.
func (h *Handlers) MapRoutes(router *gin.Engine) {
	authRoutes := router.Group("api/v1/auth")
	{
		// authRoutes.POST("/login", h.authController.Login)
		authRoutes.POST("/register", h.authController.Register)
	}
}
