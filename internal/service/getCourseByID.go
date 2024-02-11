package service

import (
	"humoAcademy/internal/models"
)

func (s *Service) GetCourseByID(courseID int) (course models.Course, err error) {


	course, err = s.Repo.GetCourseByID(courseID)

	if err != nil {
		return
	}

	return

}
