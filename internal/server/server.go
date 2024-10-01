package server

import (
	"github.com/Kenini1805/go-rest-api/config"
	"github.com/Kenini1805/go-rest-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg    *config.Config
	db     *sqlx.DB
	logger logger.Logger
	router *gin.Engine
}

func NewServer(cfg *config.Config, db *sqlx.DB, logger logger.Logger) *Server {
	// Create a new Gin router.
	router := gin.New()

	// Use Gin's default logger and recovery middleware.
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Add your routes here.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

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
