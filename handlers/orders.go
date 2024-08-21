package handlers

import (
	"net/http"

	"github.com/brunompx/angula/model"
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

func (h *Handler) HandleEditOrder(w http.ResponseWriter, r *http.Request) {
	isAddingProduct := r.URL.Query().Get("isAddingProduct") == "true"

	products, err := h.store.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order := model.Order{
		Temp: true,
	}
	h.store.CreateOrder(&order)

	views.OrderEdit(products, order, isAddingProduct).Render(r.Context(), w)
}
