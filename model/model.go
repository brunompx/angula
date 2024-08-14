package model

import "strconv"

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
}

type OrderItem struct {
}
