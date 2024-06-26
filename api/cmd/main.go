package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/SashaMelva/smart_filter/internal/app"
	"github.com/SashaMelva/smart_filter/internal/config"
	"github.com/SashaMelva/smart_filter/internal/logger"
	"github.com/SashaMelva/smart_filter/internal/memory/connection"
	storage "github.com/SashaMelva/smart_filter/internal/memory/storage/postgre"
	"github.com/SashaMelva/smart_filter/internal/server/http"
)

// @title Todo App API
// @version 1.0
// @description API Server for smart filter

// @host localhost:8082
// @BasePath /

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization

func main() {
	configFile := "../configs/"
	config := config.New(configFile)
	log := logger.New(config.Logger, "../logs/")

	connectionDB := connection.New(config.DataBase, log)

	memstorage := storage.New(connectionDB.StorageDb, log)
	app := app.New(log, memstorage, config.Tokens)

	httpServer := http.NewServer(log, app, config.HttpServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		httpServer.Stop(ctx)
	}()

	log.Info("Services is running...")
	log.Debug("Debug mode enabled")

	httpServer.Start(ctx)
}
