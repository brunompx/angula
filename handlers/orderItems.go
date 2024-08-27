package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/brunompx/angula/components"
	"github.com/brunompx/angula/model"
)

func (h *Handler) HandleAddOrderItem(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
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
	order, err := h.store.FindTempOrder()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	addOrderItem(&order, product)

	go h.store.UpdateOrder(&order)
	//log.Printf("Binomial took %s", time.Since(start))
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

	//fmt.Println("deleteOrderItem IO------------------: ", productID)
	//for _, o := range order.OrderItems {
	//	fmt.Println("deleteOrderItem PRE productID: ", o.ProductID)
	//	fmt.Println("deleteOrderItem PRE  Quantity: ", o.Quantity)
	//}

	shouldDelete, orderItem := deleteOrderItem(&order, productID)

	if shouldDelete {
		go h.store.DeleteOrderItem(&orderItem)
	} else {
		go h.store.UpdateOrder(&order)
	}

	components.OrderItemsPanel(order.OrderItems).Render(r.Context(), w)
}

func deleteOrderItem(order *model.Order, productID int) (bool, model.OrderItem) {
	shouldDelete := false
	orderItem := model.OrderItem{}
	for i, item := range order.OrderItems {
		if item.ProductID == productID {
			oi := &order.OrderItems[i]
			if oi.Quantity == 1 {
				shouldDelete = true
				orderItem = item
				order.OrderItems = slices.Delete(order.OrderItems, i, i+1)
			} else {
				oi.Quantity -= 1
				oi.PriceTotal = oi.PriceTotal - oi.Price
			}
		}
	}
	return shouldDelete, orderItem
}
