package repository

import (
	"context"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetCourseByID(courseID int) (course models.Course, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
	SELECT
		id, 
		name, 
		start_date, 
		duration,
		schedule,
		age_limit,
		registration_end_date,
		address,
		description,
		mentor,
		format,
		language
	FROM 
		course 
	WHERE 
		registration_end_date > now() 
	AND
		id = $1`, courseID)

	err = row.Scan(&course.ID)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}

		repo.Logger.WithFields(logrus.Fields{
			"course_id": courseID,
			"err":       err,
		}).Error("error in repo, GetCourseByID")
	}

	return
}
