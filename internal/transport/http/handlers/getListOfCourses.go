package handlers

import (
	"humoAcademy/pkg/response"
	"net/http"
)

func (h *Handler) GetListOfCourses(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)


	courses, err := h.svc.GetListOfCourses()

	if err != nil {
		resp = response.InternalServer
		return
	}


	resp = response.Success
	resp.Payload = courses
}
