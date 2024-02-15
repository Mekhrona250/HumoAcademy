package repository

import (
	"context"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetAllApplicationsByCourseID(courseID int) (applications []models.Application, err error) {
	rows, err := repo.Conn.Query(context.Background(), `
			SELECT
				id,
				user_id
			FROM 
				application 
			WHERE 
				course_id = $1 `, courseID)

	applications = make([]models.Application, 0)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	defer rows.Close()

	for rows.Next() {
		var application models.Application

		err := rows.Scan(&application.ID, &application.UserId)

		if err != nil {
			log.Fatal(err)
		}

		applications = append(applications, application)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}

		repo.Logger.WithFields(logrus.Fields{
			"course_id": courseID,
			"err": err,
		}).Error("error in repo, GetAllApplicationsByCourseID")
	}

	return
}
