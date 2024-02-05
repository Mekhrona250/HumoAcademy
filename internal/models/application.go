package models

type Application struct {
	Id       int `json:"id"`
	UserId   int `json:"user_id"`
	CourseId int `json:"course_id"`
}
