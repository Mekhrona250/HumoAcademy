package repository

import (
	"context"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (repo *Repository) GetCourseById(courseId int) (course models.Course, err error) {
	rows, err := repo.Conn.Query(context.Background(), `
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
				id = $1 `, courseId) 

	err = rows.Scan(
		&course.Id,
		&course.Name,
		&course.StartDate,
		&course.Duration,
		&course.Schedule,
		&course.AgeLimit,
		&course.Address,
		&course.Description,
		&course.Mentor,
		&course.Format,
		&course.Language,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrDataNotFound
		}

		repo.Logger.WithFields(logrus.Fields{
			"course_id": courseId,
			"err":   err,
		}).Error("error in repo, GetCourseById")
	}

	return
}
