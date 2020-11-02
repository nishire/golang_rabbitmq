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
	// var ratePlanData model.RatePlan
	// var roomData model.Room

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
	go func() {
		for message := range messages {
			unmarshalError := json.Unmarshal(message.Body, &offerData)
			if unmarshalError != nil {
				log.Fatal("Unmarshal Offer Failed: ", unmarshalError)
			}
			for index, _ := range offerData.Offers {
				fmt.Println("Message: ", offerData.Offers[index].Hotel)

				hotelData = offerData.Offers[index].Hotel
				connnn := DbConn()
				fmt.Println("::::::::::::::::::::::::::::", connnn)
				_, err := connnn.Prepare("INSERT INTO rate_plan(hotel_id, rate_plan_id) VALUES(123,123)")
				if err != nil {
					panic(err.Error())
				}
				// ConnectToDB()
				// result := Db.Create(&hotelData)
				// DB, err := gorm.Open("mysql", DbURL(BuildDBConfig()))
				// if err != nil {
				// 	fmt.Println("Status:", err)
				// }
				// defer DB.Close()
				// DB.AutoMigrate(&model.Hotel{})

				// createError := CreateHotel(&hotelData)
				// if err != nil {
				// 	log.Fatal("Create Hotel Failed: ", createError)
				// }
			}
		}
	}()
	log.Info("Successfully Connected To RabbitMQ Instance")
	log.Info("Waiting For Further Messages")
	<-forever
}
