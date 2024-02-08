package repository

import (
	"context"
	"humoAcademy/internal/models"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateUser(user models.User) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO users(
			phone_number, 
			password,
			name,
			surname,
			date_of_birth,
			role_id
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			2
		)
	`, user.PhoneNumber, user.Password, user.Name, user.Surname, user.DateOfBirth)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user": user,
			"err":  err,
		}).Error("error in repo, CreateUser")
	}

	return
}
