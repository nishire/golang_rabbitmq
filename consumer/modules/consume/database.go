package consume

import (
	"github.com/nishire/golang_rabbitmq/consumer/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/assignment?charset=utf8mb4&parseTime=True&loc=Local"
	db, connectErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if connectErr != nil {
		log.Fatal("Connection To SQL Failed: ", connectErr)
	}
	return db
}

func InsertHotelRow(dbConn *gorm.DB, objectData model.Hotel) {

	amenities := objectData.Amenities
	for _, amen := range amenities {
		var temp model.Amenity
		temp.HotelId = objectData.HotelId
		temp.Amenity = amen

		dbConn.Table("amenities")
		result := dbConn.Create(&temp)
		if result.Error != nil {
			log.Error("Insert Hotel Data Failed: ", result.Error)
		}
	}

	var hotelData model.HotelTable
	hotelData.HotelId = objectData.HotelId
	hotelData.Name = objectData.Name
	hotelData.Country = objectData.Country
	hotelData.Address = objectData.Address
	hotelData.Latitude = objectData.Latitude
	hotelData.Longitude = objectData.Longitude
	hotelData.Telephone = objectData.Telephone
	hotelData.Description = objectData.Description
	hotelData.RoomCount = objectData.RoomCount
	hotelData.Currency = objectData.Currency

	dbConn.Table("hotel")
	result := dbConn.Create(&hotelData)
	if result.Error != nil {
		log.Error("Insert Hotel Data Failed: ", result.Error)
	}
}

func InsertRoomRow(dbConn *gorm.DB, objectData model.Room) {
	var temp model.CapacityTable
	temp.HotelId = objectData.HotelId
	temp.MaxAdults = objectData.Capacity.MaxAdults
	temp.ExtraChildren = objectData.Capacity.ExtraChildren

	dbConn.Table("capacity")
	result := dbConn.Create(&temp)
	if result.Error != nil {
		log.Error("Insert Capacity Data Failed: ", result.Error)
	}

	var roomData model.RoomTable
	roomData.HotelId = objectData.HotelId
	roomData.RoomId = objectData.RoomId
	roomData.Name = objectData.Name
	roomData.Description = objectData.Description

	dbConn.Table("room")
	result = dbConn.Create(&roomData)
	if result.Error != nil {
		log.Error("Insert Room Data Failed: ", result.Error)
	}
}

func InsertRatePlanRow(dbConn *gorm.DB, objectData model.RatePlan) {
	otherConditions := objectData.OtherConditions
	for _, cond := range otherConditions {
		var temp model.OtherConditions
		temp.HotelId = objectData.HotelId
		temp.Condition = cond

		dbConn.Table("other_conditions")
		result := dbConn.Create(&temp)
		if result.Error != nil {
			log.Error("Insert Rate Data Failed: ", result.Error)
		}
	}

	for _, cancel := range objectData.CancellationPolicy {
		var temp model.CancellationPolicyTable
		temp.HotelId = objectData.HotelId
		temp.Type = cancel.Type
		temp.ExpiresDaysBefore = cancel.ExpiresDaysBefore

		dbConn.Table("cancellation_policy")
		result := dbConn.Create(&temp)
		if result.Error != nil {
			log.Error("Insert Rate Data Failed: ", result.Error)
		}
	}

	var roomPlanData model.RatePlanTable
	roomPlanData.HotelId = objectData.HotelId
	roomPlanData.RatePlanId = objectData.RatePlanId
	roomPlanData.Name = objectData.Name
	roomPlanData.MealPlan = objectData.MealPlan

	dbConn.Table("rate_plan")
	result := dbConn.Create(&roomPlanData)
	if result.Error != nil {
		log.Error("Insert Rate Plan Data Failed: ", result.Error)
	}
}
