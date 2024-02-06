package handlers

import (
	"encoding/json"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"humoAcademy/pkg/response"
	"net/http"
)

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.Course

	err := json.NewDecoder(r.Body).Decode(&inputData)

	if err != nil {
		resp = response.BadRequest
		return
	}

	
	if inputData.Duration < 1 {
		resp = response.BadRequest
		return
	}


	err = h.svc.CreateCourse(inputData)

	if err != nil {
		if err == errors.ErrAlreadyHasUser {
			resp.Code = 409
			resp.Message = err.Error()
			return
		}

		resp = response.InternalServer
		return
	}

	resp = response.Success
}
