package repository

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) DeleteCourseByID(courseID int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
			DELETE FROM course 
			WHERE id = $1`, courseID)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"course_id": courseID,
			"err":       err,
		}).Error("error in repo, CreateCourse")
	}

	return
}
