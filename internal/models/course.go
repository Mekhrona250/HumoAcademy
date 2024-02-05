package models

import "time"

type Course struct {
	Id                  int       `json:"id"`
	Name                string    `json:"name"`
	StartDate           time.Time `json:"start_date"`
	Duration            int       `json:"duration"`
	Schedule            string    `json:"schedule"`
	AgeLimit            int       `json:"age_limit"`
	RegistrationEndDate time.Time `json:"registration_end_date"`
	Address             string    `json:"address"`
	Description         string    `json:"description"`
	Mentor              string    `json:"mentor"`
	Format              string    `json:"format"`
	Language            string    `json:"language"`
}
