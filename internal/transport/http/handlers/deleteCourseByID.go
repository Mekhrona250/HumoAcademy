package handlers

import (
	"humoAcademy/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
)

func (h *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	course := r.URL.Query().Get("courseID")

	userID, ok := context.Get(r, "userID").(int)

	if !ok {
		resp = response.InternalServer
		return
	}

	courseID, err := strconv.Atoi(course)
	if err != nil {
		resp = response.BadRequest
	}

	err = h.svc.DeleteCourseByID(courseID, int(userID))

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
