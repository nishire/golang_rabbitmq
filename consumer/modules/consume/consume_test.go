package consume

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nishire/golang_rabbitmq/consumer/model"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func TestConsumeMessage(t *testing.T) {
	var offerData model.Master
	var hotelData model.Hotel
	var ratePlanData model.RatePlan
	var roomData model.Room

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatal("Connection To RabbitMQ Failed: ", err)
	}
	defer connection.Close()

	t.Log("Successfully Connected To RabbitMQ Instance")

	channel, createChannelError := connection.Channel()
	if createChannelError != nil {
		t.Fatal("Channel Creation Failed: ", createChannelError)
	}

	defer channel.Close()

	messages, err := channel.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		t.Fatal("Consume From Channel Failed: ", err)
	}
	dbConn := ConnectToDB()
	go func() {
		for message := range messages {
			unmarshalError := json.Unmarshal(message.Body, &offerData)
			if unmarshalError != nil {
				t.Fatal("Unmarshal Offer Failed: ", unmarshalError)
			}
			for index, _ := range offerData.Offers {
				fmt.Println("Message: ", offerData.Offers[index].Hotel)
				hotelData = offerData.Offers[index].Hotel
				ratePlanData = offerData.Offers[index].RatePlan
				roomData = offerData.Offers[index].Room

				////////////////////////////////////Database Operations For Hotel/////////////////////
				amenities := hotelData.Amenities
				for _, amen := range amenities {
					var temp model.Amenity
					temp.HotelId = hotelData.HotelId
					temp.Amenity = amen

					dbConn.Table("amenities")
					result := dbConn.Create(&temp)
					if result.Error != nil {
						log.Error("Insert Hotel Data Failed: ", result.Error)
					}
				}

				var hotelDataTable model.HotelTable
				hotelDataTable.HotelId = hotelData.HotelId
				hotelDataTable.Name = hotelData.Name
				hotelDataTable.Country = hotelData.Country
				hotelDataTable.Address = hotelData.Address
				hotelDataTable.Latitude = hotelData.Latitude
				hotelDataTable.Longitude = hotelData.Longitude
				hotelDataTable.Telephone = hotelData.Telephone
				hotelDataTable.Description = hotelData.Description
				hotelDataTable.RoomCount = hotelData.RoomCount
				hotelDataTable.Currency = hotelData.Currency

				dbConn.Table("hotel")
				result := dbConn.Create(&hotelDataTable)
				if result.Error != nil {
					t.Error("Insert Hotel Data Failed: ", result.Error)
				}

				//////////////////////////////////////////////////////Database Operations For Room//////////////////////////////////////
				var temp model.CapacityTable
				temp.HotelId = roomData.HotelId
				temp.MaxAdults = roomData.Capacity.MaxAdults
				temp.ExtraChildren = roomData.Capacity.ExtraChildren

				dbConn.Table("capacity")
				result = dbConn.Create(&temp)
				if result.Error != nil {
					log.Error("Insert Capacity Data Failed: ", result.Error)
				}

				var roomDataTable model.RoomTable
				roomDataTable.HotelId = roomData.HotelId
				roomDataTable.RoomId = roomData.RoomId
				roomDataTable.Name = roomData.Name
				roomDataTable.Description = roomData.Description

				dbConn.Table("room")
				result = dbConn.Create(&roomDataTable)
				if result.Error != nil {
					t.Error("Insert Room Data Failed: ", result.Error)
				}

				////////////////////////////////////////////////////////Database Operations For Rate///////////////////////////////////
				otherConditions := ratePlanData.OtherConditions
				for _, cond := range otherConditions {
					var temp model.OtherConditions
					temp.HotelId = ratePlanData.HotelId
					temp.Condition = cond

					dbConn.Table("other_conditions")
					result := dbConn.Create(&temp)
					if result.Error != nil {
						log.Error("Insert Rate Data Failed: ", result.Error)
					}
				}

				for _, cancel := range ratePlanData.CancellationPolicy {
					var temp model.CancellationPolicyTable
					temp.HotelId = ratePlanData.HotelId
					temp.Type = cancel.Type
					temp.ExpiresDaysBefore = cancel.ExpiresDaysBefore

					dbConn.Table("cancellation_policy")
					result := dbConn.Create(&temp)
					if result.Error != nil {
						log.Error("Insert Rate Data Failed: ", result.Error)
					}
				}

				var ratePlanDataTable model.RatePlanTable
				ratePlanDataTable.HotelId = ratePlanData.HotelId
				ratePlanDataTable.RatePlanId = ratePlanData.RatePlanId
				ratePlanDataTable.Name = ratePlanData.Name
				ratePlanDataTable.MealPlan = ratePlanData.MealPlan

				dbConn.Table("rate_plan")
				result = dbConn.Create(&ratePlanDataTable)
				if result.Error != nil {
					t.Error("Insert Rate Plan Data Failed: ", result.Error)
				}
			}

		}
	}()
	t.Log("Successfully Connected To RabbitMQ Instance")
	t.Log("Waiting For Further Messages")
}
