package repository

import (
	"context"
	"humoAcademy/internal/models"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) ChangeUser(user models.User, userID int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
			UPDATE
				users 
			SET 
			phone_number = $1, 
			name = $2,
			surname = $3,
			date_of_birth = $4,
			role_id = $5
			WHERE 
				id = $6`, user.PhoneNumber, user.Name, user.Surname, user.DateOfBirth, 2, userID)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"users":    user,
			"user_id": userID,
			"err":       err,
		}).Error("error in repo, ChangeUser")
	}

	return
}
