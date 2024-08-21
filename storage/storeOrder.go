package storage

import (
	"fmt"

	"github.com/brunompx/angula/model"
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

func (s *Storage) CreateOrder(o *model.Order) (*model.Order, error) {
	result := s.db.Create(&o)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return o, nil
}
