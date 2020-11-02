package publish

import (
	"testing"

	"github.com/streadway/amqp"
)

func TestPublishMessage(t *testing.T) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatal("Connection To RabbitMQ Failed: ", err)
	}
	defer connection.Close()

	t.Log("Successfully Connected To RabbitMQ Instance")

	channel, createChannelError := connection.Channel()
	if createChannelError != nil {
		t.Errorf("Channel Creation Failed %s: ", createChannelError)
	}

	defer channel.Close()

	queue, queueError := channel.QueueDeclare("TestQueue", false, false, false, false, nil)
	if queueError != nil {
		t.Errorf("Queue Creation Failed %s: ", queueError)
	}

	t.Log("Queue Details: ", queue)

	publishError := channel.Publish("", queue.Name, false, false,
		amqp.Publishing{
			ContentType: "plain/text",
			Body:        []byte("Nishant"),
		})

	if publishError != nil {
		t.Errorf("Publish Message Failed %s: ", publishError)
	}

	t.Log("Publish Message Successful")
}
