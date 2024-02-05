package service

import (
	"humoAkademy/internal/repository"
	"humoAkademy/pkg/config"
	"humoAkademy/internal/models"

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
}

func NewService(repo repository.RepositoryInterface, config *config.Config, logger *logrus.Logger) ServiceInterface {
	return &Service{
		Repo:   repo,
		Config: config,
		Logger: logger,
	}
}
