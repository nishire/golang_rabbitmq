package publish

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func PublishMessage(byteData []byte) {
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

	queue, queueError := channel.QueueDeclare("TestQueue", false, false, false, false, nil)
	if queueError != nil {
		log.Fatal("Queue Creation Failed: ", queueError)
	}
	// here add >> channel.Qos(1, 0, false)
	log.Info("Queue Details: ", queue)

	// byteData, marshalError := json.Marshal(myData)
	// if queueError != nil {
	// 	log.Fatal("Marshal Failed: ", marshalError)
	// }

	publishError := channel.Publish("", queue.Name, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteData,
		})

	if publishError != nil {
		log.Fatal("Publish Message Failed: ", publishError)
	}

	log.Info("Publish Message Successful")
}
