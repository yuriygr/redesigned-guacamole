package repository

import (
	"github.com/go-redis/cache/v8"
	"github.com/jmoiron/sqlx"
	"github.com/yuriygr/go-posledstvie/container"
)

// Repository - Репозиторий
type Repository struct {
	storage *sqlx.DB
	cache   *cache.Cache
}

// NewRepository - Новый репозиторий
func NewRepository(container *container.Container) *Repository {
	return &Repository{container.Storage, container.Cache}
}
