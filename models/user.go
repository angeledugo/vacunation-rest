package models

import "time"

type User struct {
	Id         string    `json:"id"`
	Name       string    `json:"name" validate:"required,min=3,max=20"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
}
