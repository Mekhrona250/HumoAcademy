package repository

import (
	"context"
	"humoAcademy/internal/models"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) ChangeCourse(course models.Course, courseID int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
			UPDATE
				course 
			SET 
				start_date = $1,
				schedule = $2,
				registration_end_date = $3,
				mentor = $4
			WHERE 
				id = $5`, course.StartDate, course.Schedule, course.RegistrationEndDate, course.Mentor, courseID)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"course": course,
			"err":       err,
		}).Error("error in repo, ChangeCourse")
	}

	return
}
