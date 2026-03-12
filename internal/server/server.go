package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/cristianrisueo/shopping-api/internal/config"
	"github.com/cristianrisueo/shopping-api/internal/utils"
)

// Server holds the application dependencies shared across handlers.
type Server struct {
	config *config.Config
	db     *gorm.DB
	logger zerolog.Logger
}

// NewServer creates a new Server with the given dependencies.
func NewServer(cfg *config.Config, db *gorm.DB, logger zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

// SetupRouter registers middleware and routes, and returns the configured engine.
func (s *Server) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", func(ctx *gin.Context) {
		utils.SuccessResponse(ctx, "API is healthy", nil)
	})

	return router
}
