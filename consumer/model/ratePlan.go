package model

import "github.com/jinzhu/gorm"

type RatePlan struct {
	HotelId            string   `json:"hotel_id"`
	RatePlanId         string   `json:"rate_plan_id"`
	CancellationPolicy []Cancel `json:"cancellation_policy"`
	Name               string   `json:"name"`
	OtherConditions    []string `json:"other_conditions"`
	MealPlan           string   `json:"meal_plan"`
}

type RatePlanTable struct {
	gorm.Model
	HotelId    string `json:"hotel_id"`
	RatePlanId string `json:"rate_plan_id"`
	Name       string `json:"name"`
	MealPlan   string `json:"meal_plan"`
}

type Cancel struct {
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}

type OtherConditions struct {
	HotelId   string `json:"hotel_id"`
	Condition string `json:"condition"`
}

type CancellationPolicyTable struct {
	HotelId           string `json:"hotel_id"`
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}

// func (rp *RatePlan) TableNameRatePlan() string {
// 	return "rate_plan"
// }
