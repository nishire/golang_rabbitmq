package consume

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nishire/golang_rabbitmq/consumer/model"
	"github.com/streadway/amqp"
)

func TestConsumeMessage(t *testing.T) {
	var offerData model.Master
	// var hotelData model.Hotel
	// var ratePlanData model.RatePlan
	// var roomData model.Room

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatal("Connection To RabbitMQ Failed: ", err)
	}
	defer connection.Close()

	t.Log("Successfully Connected To RabbitMQ Instance")

	channel, createChannelError := connection.Channel()
	if createChannelError != nil {
		t.Error("Channel Creation Failed: ", createChannelError)
	}

	defer channel.Close()

	messages, err := channel.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		t.Error("Consume From Channel Failed: ", err)
	}

	forever := make(chan bool)
	go func() {
		for message := range messages {
			unmarshalError := json.Unmarshal(message.Body, &offerData)
			if unmarshalError != nil {
				t.Error("Unmarshal Offer Failed: ", unmarshalError)
			}
			for index, _ := range offerData.Offers {
				fmt.Println("Message: ", offerData.Offers[index].Hotel)

				// hotelData = offerData.Offers[index].Hotel
				connection := DbConn()
				t.Log("Coonection", connection)
			}
		}
	}()
	t.Log("Successfully Connected To RabbitMQ Instance")
	t.Log("Waiting For Further Messages")
	<-forever
}
