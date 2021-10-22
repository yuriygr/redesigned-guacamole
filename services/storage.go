package services

import (
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yuriygr/go-loggy"
)

var (
	onceStorage sync.Once
	db          *sqlx.DB
	err         error
)

// NewStorage - Создаем экземпляр хранилища
func NewStorage(config *Config, logger *loggy.Logger) *sqlx.DB {
	onceStorage.Do(func() {
		db, err = sqlx.Connect(config.Storage.Driver, config.Storage.DSN)
		if err != nil {
			logger.Error(err)
			log.Fatalln(err)
		}
		db.SetConnMaxLifetime(time.Hour)
		db.SetMaxOpenConns(config.Storage.MaxOpenConns)
	})

	// Unsafe becouse i sleep
	return db.Unsafe()
}
