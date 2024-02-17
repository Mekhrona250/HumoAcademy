package service

import (
	"humoAcademy/pkg/errors"
	"time"
)

func (s *Service) CourseRegister(userID, courseID int) (err error) {

	user, err := s.Repo.GetUserByID(userID)

	if err != nil {
		err = errors.ErrDataNotFound
		return
	}

	course, err := s.Repo.GetCourseByID(courseID)

	if err != nil {
		return
	}

	if course.RegistrationEndDate.Unix() < time.Now().Unix() {
		err = errors.ErrDataNotFound
		return
	}
	
	if user.DateOfBirth.AddDate(course.AgeLimit, 0, 0).Unix() > time.Now().Unix() {
		err = errors.ErrAccessDenied
		return
	}

	err = s.Repo.AddApplication(userID, courseID)

	return
}
