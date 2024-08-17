package model

import (
	"strconv"
	"time"
)

type Product struct {
	ID          int    `json:"id"`
	User        string `json:"user"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
	Category    string `json:"category"`
	CreatedAt   string `json:"createdAt"`
	CategoryId  int    `json:"category_id"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) IDS() string {
	return strconv.Itoa(c.ID)
}

type Order struct {
	ID           uint        `json:"id"`
	User         string      `json:"user"`
	Name         string      `json:"name"`
	Comment      string      `json:"comment"`
	ItemsDesc    string      `json:"items_desc"`
	UpdatedAt    time.Time   `json:"updated_at"`
	CheckoutAt   time.Time   `json:"checkout_at"`
	Complete     bool        `json:"complete"`
	Paid         bool        `json:"paid"`
	Delivered    bool        `json:"delivered"`
	Cancelled    bool        `json:"cancelled"`
	Price        uint        `json:"price"`
	DeliveryTime string      `json:"delivery_time"`
	DeliveryInfo string      `json:"delivery_info"`
	OrderItems   []OrderItem `json:"order_items"`
}

type OrderItem struct {
	ID         uint `json:"id"`
	OrderID    uint `json:"order_id"`
	ProductID  uint `json:"product_id"`
	Quantity   uint `json:"quantity"`
	Price      uint `json:"price"`
	PriceTotal uint `json:"price_total"`
}
