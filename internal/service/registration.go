package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"time"
)

func (s *Service) Registration(user models.User) (err error) {
	if len(user.PhoneNumber) != 9 {
		return errors.ErrIncorrectPhoneNumber
	}

	_, err = s.Repo.GetUserByPhone(user.PhoneNumber)

	if err != errors.ErrDataNotFound {
		if err == nil {
			return errors.ErrAlreadyHasUser
		}

		return
	}

	if user.Name == "" || user.Surname == "" || user.DateOfBirth == time.Now() {
		err = errors.ErrBadRequest
		return
	}

	if user.Password == "" {
		err = errors.ErrTypePassword
		return
	}
	err = s.Repo.CreateUser(user)

	return
}
