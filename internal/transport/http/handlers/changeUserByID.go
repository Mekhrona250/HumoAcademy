package handlers

import (
	"encoding/json"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/errors"
	"humoAcademy/pkg/response"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

func (h *Handler) ChangeUser(w http.ResponseWriter, r *http.Request) {
	temp := struct {
		models.User
		DateOfBirth string `json:"date_of_birth"`
	}{}

	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.User

	userID, ok := context.Get(r, "userID").(int)

	if !ok {
		resp = response.InternalServer
		return
	}

	err := json.NewDecoder(r.Body).Decode(&temp)

	if err != nil {
		resp = response.BadRequest
		return
	}

	dateOfBirth, err := time.Parse("2006-01-02", temp.DateOfBirth)
	if err != nil {
		resp = response.BadRequest
	}

	inputData = models.User{
		PhoneNumber: temp.PhoneNumber,
		Name:        temp.Name,
		Surname:     temp.Surname,
		DateOfBirth: dateOfBirth,
	}

	err = h.svc.ChangeUser(inputData, userID)

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
