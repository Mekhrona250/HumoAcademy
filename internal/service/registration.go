package service

import (
	"humoAkademy/internal/models"
	"humoAkademy/pkg/errors"
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

	err = s.Repo.CreateUser(user)

	return
}
