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
func Log(msg string, obj log.Fields, level string) {
	switch level {
	case debug:
		log.WithFields(obj).Debug(msg)
	case trace:
		log.WithFields(obj).Trace(msg)
	case warn:
		log.WithFields(obj).Warn(msg)
	case error:
		log.WithFields(obj).Error(msg)
	case fatal:
		log.WithFields(obj).Fatal(msg)
	case panic:
		log.WithFields(obj).Panic(msg)
	default:
		log.WithFields(obj).Info(msg)
	}
}
