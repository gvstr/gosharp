package log

import (
	"log"
	"os"
)

type Logger struct {
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Debug:   log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime),
		Info:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Warning: log.New(os.Stdout, "WARNING\t", log.Ldate|log.Ltime),
		Error:   log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
