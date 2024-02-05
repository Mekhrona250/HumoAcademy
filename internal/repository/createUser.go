package repository

import (
	"context"
	"humoAkademy/internal/models"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO users(
			phone_number, 
			password,
			name,
			surname,
			role_id,
			date_of_birth
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`, user.PhoneNumber, user.Password, user.Name, user.Surname, user.RoleId, user.DateOfBirth)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user": user,
			"err":  err,
		}).Error("error in repo, CreateUser")
	}

	return
}
