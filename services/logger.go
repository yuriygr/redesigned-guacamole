package services

import (
	"os"
	"sync"

	"github.com/yuriygr/go-loggy"
)

var (
	onceLogger sync.Once
	logger     *loggy.Logger
)

// NewLogger - Создает новый логгер
func NewLogger(config *Config) *loggy.Logger {
	onceLogger.Do(func() {
		flags := os.O_RDWR | os.O_CREATE

		// Create or open log file
		file, _ := os.OpenFile(config.Application.Path+config.Logger.File, flags, 0666)

		logger = loggy.NewLogger(loggy.LoggerConfig{
			Writer: file,
		})
	})
	return logger
}
