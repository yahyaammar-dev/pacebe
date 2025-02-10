package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

func Init() error {
	// Create logs directory if it doesn't exist
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Create or open log file with current date
	currentDate := time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile(
		filepath.Join("logs", fmt.Sprintf("%s.log", currentDate)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	// Initialize different loggers
	InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime)
	DebugLogger = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime)

	return nil
}

func Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	message := fmt.Sprintf(format, v...)
	InfoLogger.Printf("[%s:%d] %s", filepath.Base(file), line, message)
}

func Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	message := fmt.Sprintf(format, v...)
	ErrorLogger.Printf("[%s:%d] %s", filepath.Base(file), line, message)
}

func Debug(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	message := fmt.Sprintf(format, v...)
	DebugLogger.Printf("[%s:%d] %s", filepath.Base(file), line, message)
}
