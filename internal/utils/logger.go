package utils

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Printf("[INFO] %s: %s\n", time.Now().Format(time.RFC3339), msg)
}

func (l *Logger) Warning(msg string) {
	l.logger.Printf("[WARNING] %s: %s\n", time.Now().Format(time.RFC3339), msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Printf("[ERROR] %s: %s\n", time.Now().Format(time.RFC3339), msg)
}
