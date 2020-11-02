package consume

import (
	"encoding/json"
	"fmt"

	"github.com/nishire/golang_rabbitmq/consumer/model"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func ConsumeMessage() {
	var offerData model.Master
	var hotelData model.Hotel
	var ratePlanData model.RatePlan
	var roomData model.Room

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Connection To RabbitMQ Failed: ", err)
	}
	defer connection.Close()

	log.Info("Successfully Connected To RabbitMQ Instance")

	channel, createChannelError := connection.Channel()
	if createChannelError != nil {
		log.Fatal("Channel Creation Failed: ", createChannelError)
	}

	defer channel.Close()

	messages, err := channel.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Consume From Channel Failed: ", err)
	}

	forever := make(chan bool)

	// connecting to sql database
	dbConn := ConnectToDB()

	go func() {
		for message := range messages {
			unmarshalError := json.Unmarshal(message.Body, &offerData)
			if unmarshalError != nil {
				log.Fatal("Unmarshal Offer Failed: ", unmarshalError)
			}
			for index, _ := range offerData.Offers {
				fmt.Println("Message: ", offerData.Offers[index].Hotel)

				// extracting required data from json
				hotelData = offerData.Offers[index].Hotel
				ratePlanData = offerData.Offers[index].RatePlan
				roomData = offerData.Offers[index].Room

				// inserting data into respective tables in the form of rows
				InsertHotelRow(dbConn, hotelData)
				InsertRatePlanRow(dbConn, ratePlanData)
				InsertRoomRow(dbConn, roomData)

			}
		}
	}()
	log.Info("Successfully Connected To RabbitMQ Instance")
	log.Info("Waiting For Further Messages")
	<-forever
}
