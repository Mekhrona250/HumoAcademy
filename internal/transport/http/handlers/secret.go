package handlers

import (
	"humoAcademy/pkg/response"
	"net/http"

	"github.com/gorilla/context"
)

func (h *Handler) Secret(w http.ResponseWriter, r *http.Request) {
	var resp response.Response

	defer resp.WriteJSON(w)

	resp = response.Success

	userID, ok := context.Get(r, "userID").(int)

	if !ok {
		resp = response.InternalServer
		return
	}

	if userID == 5 {
		resp.Message = "super secret"
		return
	}

	resp.Message = "super super secret"
}
