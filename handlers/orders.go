package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/brunompx/angula/components"
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

func (h *Handler) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	ordereID, err := strconv.Atoi(r.PathValue("orderID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.store.DeleteOrder(ordereID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.HandleListOrders(w, r)
}

func (h *Handler) HandleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	ordereName := r.FormValue("name")
	DeliveryTime := r.FormValue("DeliveryTime")
	DeliveryInfo := r.FormValue("DeliveryInfo")
	Delivered := r.FormValue("Delivered") == "on"

	fmt.Println("HandleUpdateOrder name: ", ordereName)
	fmt.Println("HandleUpdateOrder DeliveryTime: ", DeliveryTime)
	fmt.Println("HandleUpdateOrder DeliveryInfo: ", DeliveryInfo)
	fmt.Println("HandleUpdateOrder Delivered: ", Delivered)

	order, err := h.store.FindTempOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//TODO: guardar cambios en el ORDER

	//h.HandleEditOrder(w, r)
	components.OrderForm(order).Render(r.Context(), w)
}
