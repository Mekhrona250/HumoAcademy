package handlers

import (
	"fmt"
	"humoAcademy/pkg/response"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/context"
)

func (h *Handler) GetListOfUsersByCourseID(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

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

	xlsx, err := h.svc.GetListOfUsersByCourseID(userID, courseID)

	if err != nil {
		resp = response.InternalServer
		return
	}

	buf, _ := xlsx.WriteToBuffer()
	http.ServeContent(w, r, "list.xlsx", time.Time{}, strings.NewReader(buf.String()))
	fmt.Println(buf)
	resp = response.Success

}
