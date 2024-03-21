package models

import "time"

type Drug struct {
	Id           string    `json:"id"`
	Name         string    `json:"name" validate:"required,min=3,max=20"`
	Approved     bool      `json:"approved"`
	Min_dose     int64     `json:"min_dose"`
	Max_dose     int64     `json:"max_dose"`
	Available_at time.Time `json:"available_at"`
}
