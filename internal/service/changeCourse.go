package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
)

func (s *Service) ChangeCourse(course models.Course, userID int) (err error) {
	user, err := s.Repo.GetUserByID(userID)

	if err != nil {
		return
	}

	if user.RoleId != 1 {
		err = errors.ErrAccessDenied
		return
	}

	err = s.Repo.ChangeCourse(course)

	return
}
