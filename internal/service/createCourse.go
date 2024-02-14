package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
)

func (s *Service) CreateCourse(course models.Course, userID int) (err error) {
	
	user, err := s.Repo.GetUserByID(userID)

	if err == errors.ErrDataNotFound {
		if err != nil {
			return errors.ErrAlreadyHasCourse
		}

		return
	}

	if user.RoleId != 1 {
		err = errors.ErrAccessDenied
		return
	}
 
	if course.Duration < 1 || course.AgeLimit < 1 {
		err = errors.ErrAccessDenied
		return
	}

	err = s.Repo.CreateCourse(course)

	return
}
