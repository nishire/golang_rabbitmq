package model

import "github.com/jinzhu/gorm"

type RatePlan struct {
	gorm.Model
	HotelId            string   `json:"hotel_id"`
	RatePlanId         string   `json:"rate_plan_id"`
	CancellationPolicy []Cancel `json:"cancellation_policy"`
	Name               string   `json:"name"`
	OtherConditions    []string `json:"other_conditions"`
	Capacity           string   `json:"meal_plan"`
}

type Cancel struct {
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}

// func (rp *RatePlan) TableNameRatePlan() string {
// 	return "rate_plan"
// }
