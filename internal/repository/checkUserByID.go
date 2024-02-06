package repository

import (
	"context"
	"humoAcademy/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) CheckUserById(userID int) (err error) {
	var idFromDB int

	row := repo.Conn.QueryRow(context.Background(), `
			select id from users 
			where id = $1`, userID)

	err = row.Scan(&idFromDB)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user_id": userID,
			"err":     err,
		}).Error("error in repo, checkUserByID")
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	return
}
