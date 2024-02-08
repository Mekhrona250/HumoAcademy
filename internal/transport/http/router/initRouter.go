package router

import (
	"humoAcademy/internal/transport/http/handlers"
	"humoAcademy/internal/transport/http/middleware"
	"humoAcademy/pkg/http"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler, mw middleware.MiddlewareInterface) *mux.Router {
	router := http.NewRouter()
	router.Use(mw.TimeDuration)

	privateRouter := router.NewRoute().Subrouter()
	// fs := net_http.FileServer(net_http.Dir("../files/"))

	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/secret", handlers.Secret).Methods("POST", "GET")

	router.HandleFunc("/api/registration", handlers.Registration).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	//clients
	router.HandleFunc("/api/courseRegister", handlers.CourseRegister).Methods("POST")

	//administrator
	router.HandleFunc("/api/createCourse", handlers.CreateCourse).Methods("POST")

	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/deleteCourse", handlers.DeleteCourse).Methods("POST")

	return router
}
