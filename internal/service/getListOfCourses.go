package service

import (
	"humoAcademy/internal/models"
)

func (s *Service) GetListOfCourses() (courses []models.Course, err error) {


	courses, err = s.Repo.GetManyCourses()

	if err != nil {
		return
	}

	return

}
