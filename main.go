package main

import (
	"github.com/nishire/golang_rabbitmq/api/modules/publisher"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("IN: Main")
	publisher.PublishMessage()
	// consumer.ConsumeMessage()
	log.Info("OUT: Main")
}
