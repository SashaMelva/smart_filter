package app

import (
	storage "github.com/SashaMelva/smart_filter/internal/memory/storage/postgre"
	"go.uber.org/zap"
)

type App struct {
	storage *storage.Storage
	log     *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, storage *storage.Storage) *App {
	return &App{
		storage: storage,
		log:     logger,
	}
}
