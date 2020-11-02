package model

import "github.com/jinzhu/gorm"

type Room struct {
	HotelId     string   `json:"hotel_id"`
	RoomId      string   `json:"room_id"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Capacity    Capacity `json:"capacity"`
}

type RoomTable struct {
	gorm.Model
	HotelId     string `json:"hotel_id"`
	RoomId      string `json:"room_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Capacity struct {
	MaxAdults     int `json:"max_adults"`
	ExtraChildren int `json:"extra_children"`
}

type CapacityTable struct {
	HotelId       string `json:"hotel_id"`
	MaxAdults     int    `json:"max_adults"`
	ExtraChildren int    `json:"extra_children"`
}
