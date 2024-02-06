package repository

import (
	"context"
	"humoAcademy/internal/models"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) CreateCourse(course models.Course) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO course(
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
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11
		)
	`, )

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"course": course,
			"err":  err,
		}).Error("error in repo, CreateCourse")
	}

	return
}
