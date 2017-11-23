package main

import (
	"io"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var config Config

func init() {
	config = getConfig()
	setupLogging(config.Service.LogOutput, config.Service.LogLevel, config.Service.LogFormat)
}

func setupLogging(output io.Writer, level log.Level, format string) {
	log.SetOutput(output)
	log.SetLevel(level)
	if strings.EqualFold(format, "json") {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	log.Info("application started")

	log.Info("application completed")
}
