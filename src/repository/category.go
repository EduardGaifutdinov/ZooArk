package repository

import (
	"errors"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
)

// CategoryRepo struct
type CategoryRepo struct{}

// NewCategoryRepo returns pointer to
// category repository with all methods
func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{}
}

// Add creates product category
// return product category or error
func (ct CategoryRepo) Add(category *domain.Category) error {
	if exist := config.DB.
		Unscoped().
		Where("name = ? AND (deleted_at >  ? OR deleted_at IS NULL)",
			category.Name, category.DeletedAt).
		Find(category).RecordNotFound(); !exist {
			return errors.New("this category already exist")
	}

	err := config.DB.Create(category).Error
	return err
}
