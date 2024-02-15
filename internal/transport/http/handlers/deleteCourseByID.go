package handlers

import (
	"humoAcademy/pkg/errors"
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

	err = h.svc.DeleteCourseByID(courseID, userID)

	if err != nil {
		if err == errors.ErrBadRequest {
			resp = response.BadRequest
			return
		} else if err == errors.ErrAccessDenied {
			resp = response.AccessDenied
			return
		}
		resp = response.InternalServer
		return
	}
	resp = response.Success
}
