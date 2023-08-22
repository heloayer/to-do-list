package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/heloayer/to-do-list/config"
	"github.com/heloayer/to-do-list/internal/controller"
	"github.com/heloayer/to-do-list/pkg/httpserver"
	"github.com/heloayer/to-do-list/pkg/logger"
	"github.com/heloayer/to-do-list/pkg/mongo"
)

func Run(cfg *config.Config) {
	logging := logger.New(cfg.Log.Level)

	mongoClient, err := mongo.New(cfg.Mongo)
	if err != nil {
		logging.Fatal(fmt.Errorf("app - fn Run - mongo.New: %w", err))
	}

	defer mongoClient.Client.Disconnect(context.Background())

	// сет UseCases для NewRouter
	handler := gin.New()
	controller.NewRouter(handler)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))
	// Ожидание сигнала
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logging.Info("app - fn Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logging.Error(fmt.Errorf("app - fn Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logging.Error(fmt.Errorf("app - fn Run - httpServer.Shutdown: %w", err))
	}

}
