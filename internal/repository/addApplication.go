package repository

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (repo *Repository) AddApplication(userID, courseID int) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO application(
			user_id,
			course_id
		) VALUES(
			$1,
			$2
		)
	`, userID, courseID)

	if err != nil {
		repo.Logger.WithFields(logrus.Fields{
			"user_id":   userID,
			"course_id": courseID,
			"err":       err,
		}).Error("error in repo, AddApplication")
	}

	return
}
