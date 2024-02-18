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

	router.HandleFunc("/api/registration", handlers.Registration).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	//clients
	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/courseRegister", handlers.CourseRegister).Methods("POST")

	router.HandleFunc("/api/getListOfCourses", handlers.GetListOfCourses).Methods("GET")

	router.HandleFunc("/api/getCourseByID", handlers.GetCourseByID).Methods("GET")

	//administrator
	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/createCourse", handlers.CreateCourse).Methods("POST")

	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/deleteCourse", handlers.DeleteCourse).Methods("POST")

	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/changeCourse", handlers.ChangeCourse).Methods("POST")

	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/changeUser", handlers.ChangeUser).Methods("POST")

	privateRouter.Use(mw.JWT)
	privateRouter.HandleFunc("/api/getListOfUsersByCourseID", handlers.GetListOfUsersByCourseID).Methods("GET")

	return router
}
