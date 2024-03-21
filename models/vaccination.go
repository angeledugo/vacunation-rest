package models

import (
	"time"
)

type Vaccination struct {
	Id      string    `json:"id"`
	Name    string    `json:"name" validate:"required,min=3,max=20"`
	Drug_id string    `json:"drug_id"`
	Dose    int64     `json:"dose"`
	Date    time.Time `json:"date"`
}
