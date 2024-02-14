package handlers

import (
	"encoding/json"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"humoAcademy/pkg/response"
	"net/http"
	"time"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	temp := struct {
		models.User
		DateOfBirth string `json:"date_of_birth"`
	}{}

	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.User

	err := json.NewDecoder(r.Body).Decode(&temp)

	if err != nil {
		resp = response.BadRequest
		return
	}

	dateOfBirth, err := time.Parse("2006-01-02",temp.DateOfBirth)
	if err != nil {
		resp = response.BadRequest
	}

	inputData = models.User{
		PhoneNumber: temp.PhoneNumber, 
		Password: temp.Password, 
		Name: temp.Name, 
		Surname: temp.Surname, 
		DateOfBirth: dateOfBirth,
	}

	err = h.svc.Registration(inputData)

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
