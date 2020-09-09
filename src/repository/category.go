package repository

import (
	"errors"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/types"
	"net/http"
	"time"
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

// Delete delete product category
// return code and error
func (ct CategoryRepo) Delete(path types.PathCategory) (int, error) {
	result := config.DB.
		Unscoped().
		Model(&domain.Category{}).
		Where("id = ? AND (deleted_at < ? OR deleted_at IS NULL)", path.CategoryID, time.Now()).
		Update("deleted_at", time.Now().UTC().Truncate(time.Hour*24).AddDate(0, 0, 1))

	if result.Error != nil {
		return http.StatusBadRequest, result.Error
	}

	if result.RowsAffected == 0 {
		return http.StatusNotFound, errors.New("category not found")
	}

	return 0, nil
}

// Get returns list of categories of passed catering ID
// returns list of categories and error
func (ct CategoryRepo) Get(date string) ([]domain.Category, int, error) {
	var categories []domain.Category

	err := config.DB.
		Unscoped().
		Where("(date = ? OR date is NULL)"+
			" AND (deleted_at > ? OR deleted_at IS NULL)"+
			" AND (deleted_at IS NULL OR date IS NULL)", date, date).
		Order("created_at").
		Find(&categories).
		Error

	return categories, 0, err
}

// GetByKey returns single category item found by key
// and error if exists
func (ct CategoryRepo) GetByKey(key, value string) (domain.Category, error) {
	var category domain.Category
	err := config.DB.
		Where(key+" = ?", value).
		First(&category).Error
	return category, err
}

// Update checks if that name already exists in provided catering
// if its exists throws and error, if not updates the reading
func (dc CategoryRepo) Update(path types.PathCategory, category *domain.Category) (int,error) {
	if categoryExist := config.DB.
		Where("name = ? AND id = ? AND id(deleted_at > ? OR deleted_at IS NULL",
			category.Name, path.CategoryID, time.Now()).
		Find(&category).
		RowsAffected; categoryExist == 0 {
			if nameExist := config.DB.
				Where("name = ?", category.Name).
				Find(&category).
				RowsAffected; nameExist != 0 {
					return http.StatusBadRequest, errors.New("category with that name already exist")
			}
	}

	if categoryNotExist := config.DB.
		Unscoped().
		Model(&domain.Category{}).
		Where("id = ? AND (deleted_at > ? OR deleted_at IS NULL)", path.CategoryID, time.Now()).
		Update(category); categoryNotExist.RowsAffected == 0 {
			return http.StatusNotFound, errors.New("category not found")
	}

	return 0, nil
}