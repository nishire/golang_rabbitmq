package model

import "github.com/jinzhu/gorm"

type Master struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	gorm.Model
	Hotel    Hotel    `json:"hotel"`
	Room     Room     `json:"room"`
	RatePlan RatePlan `json:"rate_plan"`
}
