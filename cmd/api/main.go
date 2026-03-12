package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cristianrisueo/shopping-api/internal/config"
	"github.com/cristianrisueo/shopping-api/internal/database"
	"github.com/cristianrisueo/shopping-api/internal/logger"
)

func main() {
	// Initialize logger
	log := logger.NewLogger()

	// Load server configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	gin.SetMode(config.Server.GinMode)

	// Connect to the database
	db, err := database.NewConnection(config.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database connection")
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	defer sqlDB.Close()
	log.Info().Msg("Database connection established successfully")

}
