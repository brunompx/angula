package model

import (
	"strconv"
	"time"
)

type Product struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	User        string    `gorm:"column:user;not null" json:"user"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description;not null" json:"description"`
	Active      bool      `gorm:"column:active;not null" json:"active"`
	Price       int       `gorm:"column:price;not null" json:"price"`
	Stock       int       `gorm:"column:stock;not null" json:"stock"`
	Category    string    `gorm:"column:category;not null" json:"category"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	CategoryId  int       `json:"category_id"`
}

type Category struct {
	ID   int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

func (c *Category) IDS() string {
	return strconv.Itoa(c.ID)
}

type Order struct {
	ID           int         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	User         string      `gorm:"column:user;not null" json:"user"`
	Name         string      `gorm:"column:name;not null" json:"name"`
	Comment      string      `gorm:"column:comment;not null" json:"comment"`
	ItemsDesc    string      `gorm:"column:itemsDesc;not null" json:"itemsDesc"`
	UpdatedAt    time.Time   `gorm:"column:updatedAt;not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	CheckoutAt   time.Time   `gorm:"column:checkoutAt;not null;default:CURRENT_TIMESTAMP" json:"checkoutAt"`
	Complete     bool        `gorm:"column:complete;not null" json:"complete"`
	Paid         bool        `gorm:"column:paid;not null" json:"paid"`
	Delivered    bool        `gorm:"column:delivered;not null" json:"delivered"`
	Cancelled    bool        `gorm:"column:cancelled;not null" json:"cancelled"`
	Price        int         `gorm:"column:price;not null" json:"price"`
	DeliveryTime string      `gorm:"column:delivery_time;not null" json:"deliveryTime"`
	DeliveryInfo string      `gorm:"column:delivery_info;not null" json:"deliveryInfo"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID" json:"orderItems"`
}

type OrderItem struct {
	ID         int `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	OrderID    int `gorm:"column:order_id;not null" json:"orderId"`
	ProductID  int `gorm:"column:product_id;not null" json:"productId"`
	Quantity   int `gorm:"column:quantity;not null" json:"quantity"`
	Price      int `gorm:"column:price;not null" json:"price"`
	PriceTotal int `gorm:"column:price_total;not null" json:"priceTotal"`
}
