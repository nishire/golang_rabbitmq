package main

import (
	"github.com/nishire/golang_rabbitmq/consumer/modules/consume"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("IN: Main")
	consume.ConsumeMessage()
	log.Info("OUT: Main")
}
