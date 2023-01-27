package main

import (
	"context"
	"go.uber.org/zap"
	"konntent-authentication-service/pkg/constants"
	"konntent-authentication-service/pkg/nrclient"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := initLogger()

	env := os.Getenv(constants.ConfigEnvKey)
	conf, cErr := initConfig(env)
	if cErr != nil {
		logger.Error("an error occurred on init config >>> ", zap.Error(cErr))
		return
	}

	app, err := boot(logger, conf.Application)
	if err != nil {
		logger.Fatal("Something went wrong while utilizing the server.", zap.Error(err))
	}

	registrar(logger)
	migrate(logger, app.pgInstance)

	sv := initServer(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	go log.Fatal(sv.Listen(":" + port))

	graceful(logger, sv, app.nrInstance)
}

func graceful(l *zap.Logger, a *fiber.App, nr nrclient.NewRelicInstance) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	nr.Application().Shutdown(constants.AppGracefulTimeout * time.Second)
	_, cancel := context.WithTimeout(context.Background(), constants.AppGracefulTimeout*time.Second)
	defer cancel()
	if err := a.Shutdown(); err != nil {
		l.Fatal("unexpected error on shut down", zap.Error(err))
	}
}
