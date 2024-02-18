package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
)

func (s *Service) ChangeUser(user models.User, userID int) (err error) {
	oldUser, err := s.Repo.GetUserByID(userID)

	if err != nil {
		return
	}

	if user.Name != "" {
		oldUser.Name = user.Name
	}

	if !user.DateOfBirth.IsZero() {
		oldUser.DateOfBirth = user.DateOfBirth
	}

	if user.Surname != "" {
		oldUser.Surname = user.Surname
	}

	if user.PhoneNumber != "" {
		oldUser.PhoneNumber = user.PhoneNumber
	}

	err = s.Repo.ChangeUser(oldUser, userID)

	if err != nil {
		err = errors.ErrBadRequest
		return
	}

	return
}
