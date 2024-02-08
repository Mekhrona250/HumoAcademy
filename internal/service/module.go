package service

import (
	"humoAcademy/internal/models"
	"humoAcademy/internal/repository"
	"humoAcademy/pkg/config"

	"github.com/sirupsen/logrus"
)

type Service struct {
	Repo   repository.RepositoryInterface
	Config *config.Config
	Logger *logrus.Logger
}

type ServiceInterface interface {
	Registration(user models.User) (err error)
	Login(user models.User) (accessToken string, err error)
	CheckUserById(userID int) (err error)
	CreateCourse(course models.Course, userID int) (err error)
	DeleteCourseByID(courseID, userID int) (err error)
	CourseRegister(userID, courseID int) (err error)
}

func NewService(repo repository.RepositoryInterface, config *config.Config, logger *logrus.Logger) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
		Logger: logger,
	}
}
