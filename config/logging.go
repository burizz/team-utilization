package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LoggingConfig() {
	logFormat := os.Getenv("LOG_FORMAT")
	logLevel := os.Getenv("LOG_LEVEL")
	logOutput := os.Getenv("LOG_OUTPUT")

	switch logFormat {
	case "TEXT", "TXT":
		log.SetFormatter(&log.TextFormatter{})
	case "JSON", "JSN":
		log.SetFormatter(&log.JSONFormatter{}) // log as JSON instead of the default ASCII formatter.
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	// can be any io.Writer
	// TODO: add some different ones as well
	switch logOutput {
	case "STDOUT":
		log.SetOutput(os.Stdout)
	case "STDERR":
		log.SetOutput(os.Stderr)
	default:
		log.SetOutput(os.Stdout)
	}

	switch logLevel {
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "ERR", "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "WARN", "WARNING":
		log.SetLevel(log.WarnLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "DBG", "DEBUG":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.Infof("LogLevel is set to %v", log.GetLevel())
}
