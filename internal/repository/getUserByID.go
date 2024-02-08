package repository

import (
	"context"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetUserByID(userID int) (user models.User, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
			SELECT 
				id, 
				phone_number, 
				password, 
				name, 
				surname, 
				role_id, 
				date_of_birth 
			FROM 
				users 
			WHERE 
				id = $1`, userID)

	err = row.Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.RoleId,
		&user.DateOfBirth,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}

		repo.Logger.WithFields(logrus.Fields{
			"user_id": userID,
			"err":     err,
		}).Error("error in repo, GetUserByID")
	}

	return
}
