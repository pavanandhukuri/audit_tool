package logger

import "fmt"

// Init initializes the global logger

// LogInfo logs an info message
func LogInfo(message string) {
	fmt.Printf("INFO: %s\n", message)
}

// LogError logs an error message
func LogError(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

// LogDebug logs a debug message
func LogDebug(message string) {
	fmt.Printf("DEBUG: %s\n", message)
}

// LogWarn logs a warning message
func LogWarn(message string) {
	fmt.Printf("WARN: %s\n", message)
}

// LogFatal logs a fatal message
func LogFatal(message string) {
	fmt.Printf("FATAL: %s\n", message)
}

// LogPanic logs a panic message
func LogPanic(message string) {
	fmt.Printf("PANIC: %s\n", message)
}
