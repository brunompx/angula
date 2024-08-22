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
	result := s.db.Where("temp = ?", order.Temp).First(&order)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("Error-FindOrder: ", result.Error)
	}
	return order, result.Error
}

func (s *Storage) GetTempOrder() (model.Order, error) {
	var order model.Order
	order, err := s.FindTempOrder()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		order.Temp = true
		s.CreateOrder(&order)
	}
	return order, nil
}
