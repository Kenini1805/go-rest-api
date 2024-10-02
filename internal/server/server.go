package server

import (
	"github.com/Kenini1805/go-rest-api/config"
	"github.com/Kenini1805/go-rest-api/docs"
	"github.com/Kenini1805/go-rest-api/pkg/logger"
	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	db     *gorm.DB
	logger logger.Logger
	router *gin.Engine
}

func NewServer(cfg *config.Config, db *gorm.DB, logger logger.Logger) *Server {
	// Create a new Gin router.
	router := gin.New()

	docs.SwaggerInfo.Title = "Go example REST API"
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

	// Use Gin's default logger and recovery middleware.
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// Map the routes using the handlers
	handlers := NewHandlers(db)
	handlers.MapRoutes(router)

	// Initialize and return the server.
	return &Server{
		cfg:    cfg,
		db:     db,
		logger: logger,
		router: router,
	}
}

func (s *Server) Run() error {
	// Get the address and port from config.
	addr := s.cfg.Server.Port

	s.logger.Infof("Starting Gin server at %s", addr)

	// Start the Gin server.
	return s.router.Run(addr)
}
