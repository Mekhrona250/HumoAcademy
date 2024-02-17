package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"time"
)

func (s *Service) GetCourseByID(courseID int) (course models.Course, err error) {

	course, err = s.Repo.GetCourseByID(courseID)

	if course.RegistrationEndDate.Unix() < time.Now().Unix() {
		err = errors.ErrDataNotFound
		return
	}

	if err != nil {
		return
	}
	return

}
