package app

import (
	"github.com/SashaMelva/smart_filter/internal/config"
	storage "github.com/SashaMelva/smart_filter/internal/memory/storage/postgre"
	"go.uber.org/zap"
)

type App struct {
	storage *storage.Storage
	log     *zap.SugaredLogger
	Tokens  *config.Tokens
}

func New(logger *zap.SugaredLogger, storage *storage.Storage, config *config.Tokens) *App {
	return &App{
		storage: storage,
		log:     logger,
		Tokens:  config,
	}
}
