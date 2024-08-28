package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	fmt.Println("HandleUpdateOrder -------------------------------- ")

	ordereName := r.FormValue("name")
	deliveryTime := r.FormValue("DeliveryTime")
	deliveryInfo := r.FormValue("DeliveryInfo")
	paid := r.FormValue("Paid") == "on"
	delivered := r.FormValue("Delivered") == "on"

	order, err := h.store.FindTempOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(strings.TrimSpace(ordereName)) != 0 {
		fmt.Println("HandleUpdateOrder name: ", ordereName)
		order.Name = ordereName
	}
	if len(strings.TrimSpace(deliveryTime)) != 0 {
		fmt.Println("HandleUpdateOrder DeliveryTime: ", deliveryTime)
		order.DeliveryTime = deliveryTime
	}
	if len(strings.TrimSpace(deliveryInfo)) != 0 {
		fmt.Println("HandleUpdateOrder DeliveryInfo: ", deliveryInfo)
		order.DeliveryInfo = deliveryInfo
	}
	fmt.Println("HandleUpdateOrder Paid: ", paid)
	fmt.Println("HandleUpdateOrder Delivered: ", delivered)

	order.Paid = paid
	order.Delivered = delivered

	go h.store.UpdateOrder(&order)

	components.OrderForm(order).Render(r.Context(), w)
}

func (h *Handler) HandleAddOrder(w http.ResponseWriter, r *http.Request) {
	order, err := h.store.FindTempOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(strings.TrimSpace(order.Name)) == 0 {
		components.ValidationError("Please set a Name to the order").Render(r.Context(), w)
		return
	}
	if len(order.OrderItems) < 1 {
		components.ValidationError("Add at leat one product").Render(r.Context(), w)
		return
	}
	order.Temp = false
	go h.store.UpdateOrder(&order)
	h.HandleListOrders(w, r)
}
