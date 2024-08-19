package handlers

import (
	"net/http"

	"github.com/brunompx/angula/views"
)

func (h *Handler) HandleListOrders(w http.ResponseWriter, r *http.Request) {

	orders, err := h.store.GetOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	views.Orders(orders).Render(r.Context(), w)
}
