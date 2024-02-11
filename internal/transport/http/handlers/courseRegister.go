package handlers

import (
	"humoAcademy/pkg/errors"
	"humoAcademy/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
)

func (h *Handler) CourseRegister(w http.ResponseWriter, r *http.Request) {
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
	err = h.svc.CourseRegister(userID, courseID)

	if err != nil {
		if err == errors.ErrAlreadyHasCourse {
			resp.Code = 409
			resp.Message = err.Error()
			return
		}

		resp = response.InternalServer
		return
	}

	resp = response.Success
}
