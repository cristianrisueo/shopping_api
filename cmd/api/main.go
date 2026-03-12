package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cristianrisueo/shopping-api/internal/config"
	"github.com/cristianrisueo/shopping-api/internal/database"
	"github.com/cristianrisueo/shopping-api/internal/logger"
	"github.com/cristianrisueo/shopping-api/internal/server"
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

	defer sqlDB.Close()
	log.Info().Msg("Database connection established successfully")

	// Initialize the http server in a separate goroutine
	server := server.NewServer(config, db, log)
	router := server.SetupRouter()

	httpServer := &http.Server{
		Addr:         ":" + config.Server.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	log.Info().Msgf("Server started on port %s", config.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Attempt graceful shutdown (after receiving signal to close the channel)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to gracefully shutdown server")
	}

	log.Info().Msg("Server shutdown gracefully")

}
