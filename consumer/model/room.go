package model

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	HotelId     string   `json:"hotel_id"`
	RoomId      string   `json:"room_id"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Capacity    Capacity `json:"capacity"`
}

type Capacity struct {
	MaxAdults     int `json:"max_adults"`
	ExtraChildren int `json:"extra_children"`
}
