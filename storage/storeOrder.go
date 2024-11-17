package storage

import (
	"errors"
	"fmt"

	"github.com/brunompx/angula/model"
	"gorm.io/gorm"
)

type StoreOrder interface {
	GetOrders() ([]model.Order, error)
}

func (s *Storage) GetOrders() ([]model.Order, error) {

	var orders []model.Order
	result := s.db.Find(&orders)
	if result.Error != nil {
		fmt.Println()
	}
	return orders, nil
}

func (s *Storage) CreateOrder(order *model.Order) (*model.Order, error) {
	result := s.db.Create(&order)
	if result.Error != nil {
		fmt.Println("Error-CreateOrder: ", result.Error)
	}
	return order, nil
}

func (s *Storage) FindTempOrder() (model.Order, error) {
	order := model.Order{Temp: true}
	result := s.db.Where("temp = ?", order.Temp).Preload("OrderItems").First(&order)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("Error-FindOrder: ", result.Error)
	}
	return order, result.Error

}


func (s *Storage) GetTempOrder() (model.Order, error) {
	var order model.Order
	order, err := s.FindTempOrder()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		orderItems := []model.OrderItem{}
		order.Temp = true
		order.OrderItems = orderItems
		s.CreateOrder(&order)
	}
	return order, nil
}

func (s *Storage) UpdateOrder(order *model.Order) error {

	fmt.Println("UpdateOrder Paid: ", order.Paid)
	fmt.Println("UpdateOrder Delivered: ", order.Delivered)

	//Added .Select("*") to NOT only update non-zero fields, and includes the booleans
	result := s.db.Session(&gorm.Session{FullSaveAssociations: true}).Select("*").Updates(&order)

	return result.Error
}

func (s *Storage) DeleteOrder(orderID int) error {

	// Delete all items for the order
	resultItems := s.db.Where("order_id = ?", orderID).Delete(&model.OrderItem{})
	if resultItems.Error != nil {
		fmt.Println("Error-CreateOrder: ", resultItems.Error)
	}

	// Delete the order
	result := s.db.Delete(&model.Order{}, orderID)

	return result.Error
}

func (s *Storage) DeleteOrderItem(orderItem *model.OrderItem) error {
	result := s.db.Delete(&orderItem)
	return result.Error
}
