package logger

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

var once sync.Once
var logger *log.Logger

func Logger() *log.Logger {
	once.Do(
		func() {
			logger = log.New()
		})

	return logger
}
