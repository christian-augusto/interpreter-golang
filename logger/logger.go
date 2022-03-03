package logger

import "fmt"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

// Prints log error
func (l *Logger) Error(err error) {
	fmt.Printf("[ERROR] - %v \n", err)
}

// Prints log warning
func (l *Logger) Warning(log string) {
	fmt.Printf("[WARNING] - %v \n", log)
}

// Prints log info
func (l *Logger) Info(log string) {
	fmt.Printf("[INFO] - %v \n", log)
}
