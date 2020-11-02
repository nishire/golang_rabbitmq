package main

import (
	"io/ioutil"

	"github.com/nishire/golang_rabbitmq/publisher/modules/publish"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Reading the json file...")
	var filename string
	filename = `offer.json`

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Read Failed!")
	}
	log.Info("Read Completed!")
	publish.PublishMessage(byteData)
}

// use viper to access config file
