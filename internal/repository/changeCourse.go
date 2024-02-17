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
			name = $1, 
			start_date = $2,
			duration = $3,
			schedule = $4,
			age_limit = $5,
			registration_end_date = $6,
			description = $7,
			mentor = $8,
			format = $9,
			language = $10
			WHERE 
				id = $11`, course.Name, course.StartDate, course.Duration, course.Schedule, course.AgeLimit, course.RegistrationEndDate, course.Description, course.Mentor, course.Format, course.Language, courseID)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"course":    course,
			"course_id": courseID,
			"err":       err,
		}).Error("error in repo, ChangeCourse")
	}

	return
}
