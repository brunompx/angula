package handlers

import (
	"net/http"
	"strconv"

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

	order, _ := h.store.GetTempOrder()

	views.OrderEdit(products, order, isAddingProduct).Render(r.Context(), w)
}

func (h *Handler) HandleAddOrderItem(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.URL.Query().Get("productID"))

	product, err := h.store.GetProductByID(productID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order, err := h.store.FindTempOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	oItem := model.OrderItem{}

	views.OrderEdit(products, order, isAddingProduct).Render(r.Context(), w)
}
