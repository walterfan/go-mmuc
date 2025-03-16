// logger.go
package config

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logFile *os.File // global logger

// InitLogger initializes the logger to write to both console and file
func InitLogger(logDir string, logLevel string) error {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("Failed to create log directory: %v", err)
		return err
	}
	// Get the current date and time
	currentTime := time.Now()
	date := currentTime.Format("20060102")
	pid := os.Getpid()

	// Set log level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("Invalid log level: %v", err)
		return err
	}
	logrus.SetLevel(level)

	// Generate the log file name with date and PID
	logFileName := fmt.Sprintf("%s/go-mmuc_%d_%s.log", logDir, pid, date)

	// Configure lumberjack.Logger for log rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    20,   // megabytes
		MaxAge:     1,    // days
		MaxBackups: 7,    // number of old log files to retain
		LocalTime:  true, // use local time for timestamps
		Compress:   false,
	}

	// Create a multi-writer that writes to both the console and the file
	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logrus.SetOutput(multiWriter)
	logrus.SetReportCaller(true)
	// Set the log format to JSON and include filename and line number
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			shortFile := f.File[strings.LastIndex(f.File, "/")+1:]
			return "", fmt.Sprintf("%s:%d", shortFile, f.Line)
		},
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
	})
	return nil
}

// CloseLogger closes the log file when the application exits
func CloseLogger() {
	// lumberjack.Logger does not need to be closed explicitly
	// as it handles file rotation internally
}