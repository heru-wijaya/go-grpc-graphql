package lib

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	info  = "info"
	trace = "trace"
	warn  = "warn"
	debug = "debug"
	error = "error"
	fatal = "fatal"
	panic = "panic"
)

// init for setoutput and text format either json or text based
func init() {

	var file, err = os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	//log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
}

// Log function for choosing level for logging
func Log(msg string, obj map[string]interface{}, level string) {
	logger := log.WithFields(log.Fields{
		"obj": obj,
	})
	switch level {
	case debug:
		logger.Debug(msg)
	case trace:
		logger.Trace(msg)
	case warn:
		logger.Warn(msg)
	case error:
		logger.Error(msg)
	case fatal:
		logger.Fatal(msg)
	case panic:
		logger.Panic(msg)
	default:
		logger.Info(msg)
	}
}

// logger.Log("test", map[string]interface{}{
// 	"name": name,
// }, "error")
