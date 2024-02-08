package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"humoAcademy/pkg/utils"
)

func (s *Service) Login(user models.User) (accessToken string, err error) {
	if len(user.PhoneNumber) != 9 {
		err = errors.ErrIncorrectPhoneNumber
		return
	}

	userFromDB, err := s.Repo.GetUserByPhone(user.PhoneNumber)

	if err != nil {
		return
	}

	if userFromDB.Password != user.Password {
		err = errors.ErrWrongPassword
		return
	}

	accessToken, err = utils.CreateToken(s.Config.JwtSecretKey, userFromDB.PhoneNumber, userFromDB.ID)

	return
}
