package repository

import (
	"context"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetCourseByname(age int) (course models.Course, err error) {
	rows, err := repo.Conn.Query(context.Background(), `
			SELECT
				name
			FROM 
				course 
			WHERE 
				age > age_limit `)


	courseArray := models.Course{}

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&courseArray.Name)
		if err != nil {
			if err == pgx.ErrNoRows {
				err = errors.ErrDataNotFound
			}
		}

	}


	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}

		repo.Logger.WithFields(logrus.Fields{
			"age": age,
			"err":       err,
		}).Error("error in repo, GetCourseByName")
	}

	return
}
