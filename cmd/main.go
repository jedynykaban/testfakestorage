package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	//"bitbucket.org/adtoma/omakase/repos/license"
	//"cloud.google.com/go/datastore"

	//"github.com/jedynykaban/testdb/db"
	"github.com/jedynykaban/testdb/services"

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
	log.Info("testdb started")

	//mainEx()
	//implTestElements.TestEntityReflectionFun()
	cm := services.NewTimerCacheTest()
	tr := time.NewTicker(time.Second)
	go func(cmx *services.CacheManager) {
		for t := range tr.C {
			fmt.Printf("Tick at %v, cache value: %v\n", t, cmx.ReadValue())
		}
	}(cm)
	time.Sleep(time.Second * 30)
	tr.Stop()
	fmt.Println("Ticker stopped")

	log.Info("testdb completed")
}
