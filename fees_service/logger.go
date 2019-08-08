package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// Global logger used for logging in this service.
var logger *log.Logger

// Wraps the log.TextFormatter struct but uses a custom Format function according
// to the projects logstash filter.
type logStashFormatter struct {
	log.TextFormatter
}

// Format uses a custom implementation to match the logstash filter.
// Custom format: `time - name_of_service - log_level - message`.
// TODO (ccdle12): pull service name from environment.
func (l *logStashFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(
		fmt.Sprintf("%s - fees_service - %s - %s",
			entry.Time.Format(l.TimestampFormat),
			strings.ToUpper(entry.Level.String()),
			entry.Message),
	), nil
}

// Entry point, initializes the global logger variable with the custom formatter.
func init() {
	logger = &log.Logger{
		Out:   os.Stdout,
		Level: log.DebugLevel,
		Formatter: &logStashFormatter{log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true},
		},
	}
}
