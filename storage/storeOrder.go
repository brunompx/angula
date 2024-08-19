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
