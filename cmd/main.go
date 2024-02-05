package main

import (
	"humoAkademy/internal/repository"
	"humoAkademy/internal/service"
	"humoAkademy/internal/transport/http/handlers"
	"humoAkademy/internal/transport/http/middleware"
	"humoAkademy/internal/transport/http/router"
	"humoAkademy/pkg/config"
	"humoAkademy/pkg/database"
	"humoAkademy/pkg/http"
	"humoAkademy/pkg/logger"
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
