package handlers

import (
	"net/http"

	"github.com/brunompx/go-delicatessen/views"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
