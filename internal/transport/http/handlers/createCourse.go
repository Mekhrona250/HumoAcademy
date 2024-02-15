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

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	temp := struct {
		models.Course
		StartDate           string `json:"start_date"`
		RegistrationEndDate string `json:"registration_end_date"`
	}{}
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.Course

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

	StartDate, err := time.Parse("2006-01-02", temp.StartDate)
	if err != nil {
		resp = response.BadRequest
	}

	RegistrationEndDate, err := time.Parse("2006-01-02", temp.RegistrationEndDate)
	if err != nil {
		resp = response.BadRequest
	}

	inputData = models.Course{
		Name:                temp.Name,
		StartDate:           StartDate,
		Duration:            temp.Duration,
		Schedule:            temp.Schedule,
		AgeLimit:            temp.AgeLimit,
		RegistrationEndDate: RegistrationEndDate,
		Address:             temp.Address,
		Description:         temp.Description,
		Mentor:              temp.Mentor,
		Format:              temp.Format,
		Language:            temp.Language,
	}

	err = h.svc.CreateCourse(inputData, userID)

	if err == errors.ErrBadRequest {
		resp = response.BadRequest
		return
	} else if err == errors.ErrAccessDenied {
		resp = response.AccessDenied
		return
	} else if err != nil {
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
