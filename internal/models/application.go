package models

type Application struct {
	ID       int `json:"id"`
	UserId   int `json:"user_id"`
	CourseId int `json:"course_id"`
}
