package app

import (
	"go.uber.org/zap"
)

type App struct {
	storage *storage.Storage
	log     *zap.SugaredLogger
	period  entity.Period
}

func New(logger *zap.SugaredLogger, storage *storage.Storage) *App {
	return &App{
		storage: storage,
		log:     logger,
		period: map[string]string{
			"week":   "week",
			"mounth": "mounth",
			"today":  "mounth",
			"none":   "none",
		},
	}
}
