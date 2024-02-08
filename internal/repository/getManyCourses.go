package repository

import (
	"context"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetManyCourses() (courses []models.Course, err error) {
	rows, err := repo.Conn.Query(context.Background(), `
			SELECT
				id,
				name,
				schedule,
				format
			FROM 
				course 
			WHERE 
				registration_end_date > now() `)

	courses = make([]models.Course, 0)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}
	}

	defer rows.Close()

	for rows.Next() {
		var course models.Course

		err := rows.Scan(&course.ID, &course.Name, &course.Schedule, &course.Format)

		if err != nil {
			log.Fatal(err)
		}

		courses = append(courses, course)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}

		repo.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("error in repo, GetManyCourses")
	}

	return
}
