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

func (s *Storage) GetProductByID(ID int) (model.Product, error) {
	product := model.Product{ID: ID}
	result := s.db.First(&product)
	if result.Error != nil {
		fmt.Println()
	}
	return product, nil
}
