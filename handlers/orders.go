package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/brunompx/angula/components"
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

	productID, err := strconv.Atoi(r.PathValue("productID"))
	if err != nil {
		fmt.Println(err)
		return
	}

	product, err := h.store.GetProductByID(productID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("----------------------------------")
	fmt.Println("HandleAddOrderItem Product.ID: ", product.ID)

	order, err := h.store.FindTempOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	addOrderItem(&order, product)

	go h.store.UpdateOrder(&order)

	components.OrderItemsPanel(order.OrderItems).Render(r.Context(), w)
}

func addOrderItem(order *model.Order, product model.Product) {

	newItem := true
	for i, item := range order.OrderItems {

		if item.ProductID == product.ID {
			oi := &order.OrderItems[i]
			oi.Quantity += 1
			oi.PriceTotal = product.Price * oi.Quantity
			newItem = false
		}
	}
	if newItem {
		orderItem := model.OrderItem{
			ProductID:   product.ID,
			Price:       product.Price,
			PriceTotal:  product.Price,
			Quantity:    1,
			ProductName: product.Name,
		}
		ois := &order.OrderItems
		order.OrderItems = append(*ois, orderItem)
	}
}

func (h *Handler) HandleRemoveOrderItem(w http.ResponseWriter, r *http.Request) {

	productID, err := strconv.Atoi(r.PathValue("productID"))
	if err != nil {
		fmt.Println(err)
		return
	}

	order, err := h.store.FindTempOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	deleteOrderItem(&order, productID)

	go h.store.UpdateOrder(&order)

	components.OrderItemsPanel(order.OrderItems).Render(r.Context(), w)
}

func deleteOrderItem(order *model.Order, productID int) {
	for i, item := range order.OrderItems {
		if item.ProductID == productID {

			fmt.Println("deleteOrderItem IO------------------: ", productID)

			for _, o := range order.OrderItems {
				fmt.Println("deleteOrderItem PRE productID: ", o.ProductID)
				fmt.Println("deleteOrderItem PRE  Quantity: ", o.Quantity)
			}

			oi := &order.OrderItems[i]
			if oi.Quantity == 1 {
				order.OrderItems = slices.Delete(order.OrderItems, i, i+1)
			} else {
				oi.Quantity -= 1
				oi.PriceTotal = oi.PriceTotal - oi.Price
			}

			for _, o := range order.OrderItems {
				fmt.Println("deleteOrderItem POS productID: ", o.ProductID)
				fmt.Println("deleteOrderItem POS  Quantity: ", o.Quantity)
			}

		}
	}
}
