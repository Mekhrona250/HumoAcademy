package service

import (
	"humoAcademy/pkg/errors"
)

func (s *Service) DeleteCourseByID(courseID, userID int) (err error) {
	user, err := s.Repo.GetUserByID(userID)

	if err != nil {
		return
	}

	if user.RoleId != 1 {
		err = errors.ErrAccessDenied
		return
	}

	_, err = s.Repo.GetCourseByID(courseID)

	if err != nil {
		return errors.ErrBadRequest
	}

	err = s.Repo.DeleteCourseByID(courseID)

	return
}
