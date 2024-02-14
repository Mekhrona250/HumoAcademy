package handlers

import (
	"encoding/json"
	"humoAcademy/internal/models"
	"humoAcademy/pkg/response"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/context"
)

func (h *Handler) ChangeCourse(w http.ResponseWriter, r *http.Request) {
	temp := struct {
		models.Course
		StartDate           string `json:"start_date"`
		RegistrationEndDate string `json:"registration_end_date"`
	}{}
	var resp response.Response

	defer resp.WriteJSON(w)

	var inputData models.Course

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

	err = json.NewDecoder(r.Body).Decode(&temp)

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
		ID:                  temp.ID,
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

	err = h.svc.ChangeCourse(inputData, userID, courseID)

	if err != nil {
		log.Println(err)
		resp = response.InternalServer
		return
	}

	resp = response.Success
}
