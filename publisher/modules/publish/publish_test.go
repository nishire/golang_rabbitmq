package publish

import (
	"io/ioutil"
	"testing"

	"github.com/streadway/amqp"
)

func TestPublishMessage(t *testing.T) {
	t.Log("Reading the json file...")
	var filename string
	filename = `C:/Users/nishanth/Desktop/Golang_Learning/golang_rabbitmq/publisher/offer.json`

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal("Read Failed!")
	}
	t.Log("Read Completed!")

	t.Log("Connecting To RabbitMQ Instance...")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatal("Connection To RabbitMQ Failed: ", err)
	}
	defer connection.Close()

	t.Log("Successfully Connected To RabbitMQ Instance")

	channel, createChannelError := connection.Channel()
	if createChannelError != nil {
		t.Fatal("Channel Creation Failed :", createChannelError)
	}

	defer channel.Close()

	queue, queueError := channel.QueueDeclare("TestQueue", false, false, false, false, nil)
	if queueError != nil {
		t.Fatal("Queue Creation Failed :", queueError)
	}

	t.Log("Queue Details: ", queue)

	publishError := channel.Publish("", queue.Name, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteData,
		})

	if publishError != nil {
		t.Fatal("Publish Message Failed :", publishError)
	}

	t.Log("Publish Message Successful")
}
