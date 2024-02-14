package handlers

import (
	"humoAcademy/pkg/response"
	"net/http"
	"strconv"
)

func (h *Handler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	course := r.URL.Query().Get("courseID")

	courseID, err := strconv.Atoi(course)
	if err != nil {
		resp = response.BadRequest
	}

	courseByID, err := h.svc.GetCourseByID(courseID)

	if err != nil {
		resp = response.NotFound
		return
	}

	resp = response.Success
	resp.Payload = courseByID
}
