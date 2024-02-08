package handlers

import (
	"encoding/json"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/response"
	"net/http"
)

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.Course
	var userID int


	err := json.NewDecoder(r.Body).Decode(&inputData)

	if err != nil {
		resp = response.BadRequest
		return
	}

 
	err = h.svc.CreateCourse(inputData, userID)

	if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
