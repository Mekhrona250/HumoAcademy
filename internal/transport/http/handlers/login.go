package handlers

import (
	"encoding/json"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"humoAcademy/pkg/response"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.User

	err := json.NewDecoder(r.Body).Decode(&inputData)

	if err != nil {
		resp = response.BadRequest

		return
	}

	token, err := h.svc.Login(inputData)

	if err == errors.ErrDataNotFound {
		resp = response.NotFound
		return
	} else if err == errors.ErrWrongPassword {
		resp.Code = 403
		resp.Message = "Wrong Password"
		return
	} else if err == errors.ErrIncorrectPhoneNumber {
		resp.Code = 400
		resp.Message = "Incorrect phone number"
		return
	} else if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
	resp.Payload = token
}
