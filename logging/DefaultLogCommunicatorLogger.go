package logging

import (
	"errors"
	"log"
)

// DefaultLogCommunicatorLogger adapts a log.Logger to the CommunicatorLogger interface
type DefaultLogCommunicatorLogger struct {
	goLog *log.Logger
}

// Log the specified message
func (dl DefaultLogCommunicatorLogger) Log(message string) {
	dl.goLog.Println(message)
}

// LogError logs the error with a message
func (dl DefaultLogCommunicatorLogger) LogError(message string, err error) {
	dl.goLog.Println(message, err)
}

// NewDefaultLogCommunicatorLogger creates a DefaultLogCommunicatorLogger with the given log.Logger
func NewDefaultLogCommunicatorLogger(goLog *log.Logger) (*DefaultLogCommunicatorLogger, error) {
	if goLog == nil {
		return nil, errors.New("goLog is required")
	}

	return &DefaultLogCommunicatorLogger{goLog}, nil
}
