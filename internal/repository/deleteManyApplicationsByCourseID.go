package repository

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) DeleteManyApplicationsByCourseID(courseID int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
			DELETE FROM application 
			WHERE course_id = $1`, courseID)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"course_id": courseID,
			"err":       err,
		}).Error("error in repo, DeleteManyApplicationsByCourseID")
	}

	return
}
