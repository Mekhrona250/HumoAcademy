package handlers

// import (
// 	"humoAcademy/pkg/response"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func (h *Handler) GetListOfUsersByCourseID(w http.ResponseWriter, r *http.Request) {
// 	var resp response.Response

// 	defer resp.WriteJSON(w)

// 	course := r.URL.Query().Get("courseID")

// 	courseID, err := strconv.Atoi(course)
// 	if err != nil {
// 		resp = response.BadRequest
// 	}

// 	applications, err := h.svc.GetListOfUsersByCourseID(courseID)

// 	if err != nil {
// 		resp = response.InternalServer
// 		return
// 	}

// 	buf, _ := xlsx.WriteToBuffer()
// 	http.ServeContent(w, r, "test.xlsx", time.Time{}, strings.NewReader(buf.String()))
// 	resp = response.Success

//}
