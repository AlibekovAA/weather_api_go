package internal

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &Logger{log.New(file, "", log.LstdFlags)}, nil
}
