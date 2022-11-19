package slogger

import "fmt"

type SLogger struct{}

func New() *SLogger {
	return &SLogger{}
}

// LogInfo logs an info message.
func (*SLogger) LogInfo(message string) {
	fmt.Println(message)
}

// LogError logs an error message.
func (*SLogger) LogError(message string, err error) {
	fmt.Printf("%s %v\n", message, err)
}
