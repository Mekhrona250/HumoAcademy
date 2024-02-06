package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
)

func (s *Service) CreateCourse(course models.Course) (err error) {
	var user models.User

	if user.RoleId != 1 {
		err = errors.ErrAccessDenied
		return
	}

	err = s.Repo.CreateCourse(course)

	return
}
