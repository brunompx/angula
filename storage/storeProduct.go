package storage

import (
	"fmt"

	"github.com/brunompx/angula/model"
)

type StoreProduct interface {
	GetProducts() ([]model.Product, error)
	GetCategories() ([]model.Category, error)
	CreateProduct(product *model.Product) (*model.Product, error)
	DeleteProduct(id string) error
	FindProduct(search string) ([]model.Product, error)
}

func (s *Storage) GetProducts() ([]model.Product, error) {
	var products []model.Product
	result := s.db.Find(&products)
	if result.Error != nil {
		fmt.Println()
	}
	return products, nil
}

func (s *Storage) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	result := s.db.Find(&categories)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return categories, nil
}

func (s *Storage) CreateProduct(p *model.Product) (*model.Product, error) {
	result := s.db.Create(&p)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return p, nil
}

func (s *Storage) DeleteProduct(id string) error {
	err := s.db.Delete(&model.Product{}, id).Error
	return err
}

func (s *Storage) FindProduct(search string) ([]model.Product, error) {
	products := []model.Product{}
	result := s.db.Where("name LIKE ? OR description LIKE ? OR category LIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&products)
	if result.Error != nil {
		fmt.Println("Error-FindProduct: ", result.Error)
	}
	return products, nil
}

/*
func (s *Storage) GetProducts() ([]model.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []model.Product
	for rows.Next() {
		product, err := scanProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *Storage) GetCategories() ([]model.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (s *Storage) CreateProduct(p *model.Product) (*model.Product, error) {

	row, err := s.db.Exec("INSERT INTO products (user,name,description,active,price,stock,category,category_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		p.User, p.Name, p.Description, p.Active, p.Price, p.Stock, p.Category, p.CategoryId)
	if err != nil {
		return nil, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	p.ID = int(id)
	return p, nil
}

func (s *Storage) DeleteProduct(id string) error {
	result, err := s.db.Exec("DELETE FROM products WHERE id = ?", id)
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return ErrorNotFound
	}
	return err
}

func (s *Storage) FindProduct(search string) ([]model.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE name LIKE ? OR description LIKE ? OR category LIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		product, err := scanProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func scanProduct(row *sql.Rows) (model.Product, error) {
	var product model.Product
	err := row.Scan(&product.ID, &product.User, &product.Name, &product.Description,
		&product.Active, &product.Price, &product.Stock, &product.Category, &product.CreatedAt, &product.CategoryId)
	return product, err
}
*/
