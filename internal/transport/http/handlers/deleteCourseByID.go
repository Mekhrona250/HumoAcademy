package handlers

import (
	"fmt"
	"humoAcademy/pkg/response"
	"net/http"

	"github.com/gorilla/context"
)

func (h *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	courseID := r.URL.Query().Get("courseID")

	userID, ok := context.Get(r, "userID").(float64)

	fmt.Println(context.Get(r, "userID"))
	if !ok {
		resp = response.InternalServer
		fmt.Println(ok)
		return
	}

	fmt.Println(courseID, userID)
}
