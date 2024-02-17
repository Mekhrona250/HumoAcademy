package handlers

import (
	"humoAcademy/pkg/response"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/context"
)

func (h *Handler) GetListOfUsersByCourseID(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	course := r.URL.Query().Get("courseID")

	userID, ok := context.Get(r, "userID").(int)

	if !ok {
		resp = response.InternalServer
		resp.WriteJSON(w)
		return
	}

	courseID, err := strconv.Atoi(course)
	if err != nil {
		resp = response.BadRequest
		resp.WriteJSON(w)
		return
	}

	xlsx, err := h.svc.GetListOfUsersByCourseID(userID, courseID)

	if err != nil {
		resp = response.InternalServer
		resp.WriteJSON(w)
		return
	}

	buf, err := xlsx.WriteToBuffer()
	if err != nil {
		resp = response.InternalServer
		resp.WriteJSON(w)
		return
	}
	http.ServeContent(w, r, "list.xlsx", time.Time{}, strings.NewReader(buf.String()))
}
