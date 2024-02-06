package handlers

import (
	"humoAcademy/internal/service"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	svc    service.ServiceInterface
	logger *logrus.Logger
}

func NewHandler(svc service.ServiceInterface, logger *logrus.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}
