package helper

import (
	log "github.com/sirupsen/logrus"
	"runtime"
	"time"
)

func Log(error bool, message string, endpoint string, userID string, statusCode int) {
	_, file, line, _ := runtime.Caller(1)

	var logEntry *log.Entry
	logFields := log.Fields{
		"status":     getStatus(error),
		"endpoint":   endpoint,
		"message":    message,
		"userID":     userID,
		"statusCode": statusCode,
		"caller":     file + ":" + string(rune(line)),
		"timestamp":  time.Now().Format("2006-01-02 15:04:05"),
	}

	if error {
		logEntry = log.WithFields(logFields)
		logEntry.Error("Error occurred")
	} else {
		logEntry = log.WithFields(logFields)
		logEntry.Info("Operation successful")
	}
}

func getStatus(error bool) string {
	if error {
		return "error"
	}
	return "info"
}
