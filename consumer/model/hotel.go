package model

import "github.com/jinzhu/gorm"

type Hotel struct {
	gorm.Model
	HotelId     string   `json:"hotel_id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Address     string   `json:"address"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Telephone   string   `json:"telephone"`
	Amenities   []string `json:"amenities"`
	Description string   `json:"description"`
	RoomCount   int      `json:"room_count"`
	Currency    string   `json:"currency"`
}
