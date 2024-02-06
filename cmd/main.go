package main

import (
	"humoAcademy/internal/repository"
	"humoAcademy/internal/service"
	"humoAcademy/internal/transport/http/handlers"
	"humoAcademy/internal/transport/http/middleware"
	"humoAcademy/internal/transport/http/router"
	"humoAcademy/pkg/config"
	"humoAcademy/pkg/database"
	"humoAcademy/pkg/http"
	"humoAcademy/pkg/logger"
)

func main() {
	conf := config.NewConfig()

	logger := logger.NewLogger(conf)

	db := database.NewDatabase(conf, logger)
	repo := repository.NewRepository(db, logger)
	svc := service.NewService(repo, conf, logger)

	handlers := handlers.NewHandler(svc, logger)

	middle := middleware.NewMiddleware(conf, svc)

	router := router.InitRouter(handlers, middle)

	server := http.NewServer(conf.ServerAddress, conf.ServerPort, router)

	server.Run()
}
