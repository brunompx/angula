package storage

import (
	"database/sql"

	"github.com/brunompx/angula/model"
)

type StoreOrder interface {
	GetOrders() ([]model.Order, error)
}

func (s *Storage) GetOrders() ([]model.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []model.Order
	for rows.Next() {
		order, err := scanOrder(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func scanOrder(row *sql.Rows) (model.Order, error) {
	var order model.Order
	err := row.Scan(&order.ID, &order.User, &order.Name, &order.Description,
		&order.Active, &order.Price, &order.Stock, &order.Category, &order.CreatedAt, &order.CategoryId)
	return order, err
}
