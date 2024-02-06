package repository

import (
	"context"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetUserByPhone(phone string) (user models.User, err error) {
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
				phone_number = $1`, phone)

	err = row.Scan(
		&user.Id, 
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
			"phone": phone,
			"err":   err,
		}).Error("error in repo, GetUserByPhone")
	}

	return
}
