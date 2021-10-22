package container

import (
	"os"

	"github.com/go-redis/cache/v8"
	"github.com/jmoiron/sqlx"
	"github.com/yuriygr/go-loggy"
	"github.com/yuriygr/go-posledstvie/services"
)

// Container - Структура с зависимостями.
// Знаю, знаю. Но не я такой, жизнь такая.
type Container struct {
	Config  *services.Config
	Session *services.Session

	Cache   *cache.Cache
	Logger  *loggy.Logger
	Storage *sqlx.DB
}

// NewContainer - Собираем зависимости
func NewContainer() *Container {
	config := services.NewConfig(os.Getenv("GORKI_PATH"))

	logger := services.NewLogger(config)
	cache := services.NewCache(config)
	storage := services.NewStorage(config, logger)
	session := services.NewSession(config, logger)

	return &Container{config, session, cache, logger, storage}
}
