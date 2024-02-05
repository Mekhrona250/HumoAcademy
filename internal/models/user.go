package models

import "time"

type User struct {
	Id          int       `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	RoleId      int       `json:"role_id"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
