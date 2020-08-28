package repository

import (
	"errors"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
)

// ProductRepo struct
type ProductRepo struct {}

//NewProductRepo returns pointer to product repository
// with all methods
func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

// Add creates new product entity
// returns error or nil
func (p ProductRepo) Add(product *domain.Product) error {
	if productExist := config.DB.
		Where("name = ?", product.Name).
		Find(product).
		RecordNotFound(); !productExist {
			return errors.New("this product already exist")
	}

	if err := config.DB.Create(product).Error; err != nil {
		return err
	}

	return nil
}