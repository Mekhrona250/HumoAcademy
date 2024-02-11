package repository

import (
	"humoAcademy/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Conn   *pgx.Conn
	Logger *logrus.Logger
}

type RepositoryInterface interface {
	CreateUser(user models.User) (err error)
	GetUserByPhone(phone string) (user models.User, err error)
	GetUserByID(userID int) (user models.User, err error)
	CreateCourse(course models.Course) (err error)
	CheckUserById(userID int) (err error)
	GetManyCourses() (courses []models.Course, err error)
	DeleteCourseByID(courseID int) (err error)
	AddApplication(userID , courseID int) (err error) 
	GetCourseByID(courseID int) (course models.Course, err error)
	ChangeCourse(course models.Course) (err error)
}

func NewRepository(conn *pgx.Conn, logger *logrus.Logger) RepositoryInterface {
	return &Repository{
		Conn:   conn,
		Logger: logger,
	}
}
