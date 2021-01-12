package lib

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// init for setoutput and text format either json or text based
func init() {

	var file, err = os.OpenFile("/tmp/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	//log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})
}

func Info(msg string, obj log.Fields) {
	log.WithFields(obj).Info(msg)
}
