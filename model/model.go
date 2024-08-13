package model

import "strconv"

type Car struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
	Make      string `json:"make"`
	Model     string `json:"model"`
	Year      string `json:"year"`
	ImageURL  string `json:"imageURL"`
	CreatedAt string `json:"createdAt"`
}

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
