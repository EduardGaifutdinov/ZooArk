package repository

import (
	"errors"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/jinzhu/gorm"
	"net/http"
)

// ClothesRepo struct
type ClothesRepo struct{}

//NewClothesRepo returns pointer to product repository
// with all methods
func NewClothesRepo() *ClothesRepo {
	return &ClothesRepo{}
}

// Add creates new product entity
// returns error or nil
func (p ClothesRepo) Add(clothes *domain.Clothes) error {
	count := clothes.Count
	thisClothes := *clothes

	if clothes.Count <= 0 || clothes.Price <= 0 {
		return errors.New("Can't add product with price or count less than 0 ")
	}

	if productExist := config.DB.
		Where("name = ?", thisClothes.Name).
		Find(clothes).
		RecordNotFound(); !productExist {
		if productExist = config.DB.
			Where("name = ? AND color = ? AND stock = ?", thisClothes.Name, thisClothes.Color, thisClothes.Stock).
			Find(&thisClothes).
			RecordNotFound(); productExist {
			if err := config.DB.Create(&thisClothes).Error; err != nil {
				return err
			}
		}
		if err := config.DB.
			Model(&thisClothes).
			Where("name = ? AND color = ? AND stock = ?", &thisClothes.Name, &thisClothes.Color, &thisClothes.Stock).
			Update("count", gorm.Expr("count + ?", count)).
			Error; err != nil {
			return err
		}
		return nil
	}

	if err := config.DB.Create(clothes).Error; err != nil {
		return err
	}

	return nil
}

// GetByKey return single product item found by key
// and error if exist
func (p ClothesRepo) GetByKey(key, value string) (domain.Clothes, error) {
	var product domain.Clothes
	err := config.DB.
		Where("? = ?", key, value).
		First(&product).Error
	return product, err
}

// Get return list of clothes or error
func (p ClothesRepo) Get() ([]domain.Clothes, int, error) {
	var clothes []domain.Clothes

	if err := config.DB.
		Order("created_at").
		Find(&clothes).
		Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return []domain.Clothes{}, http.StatusBadRequest, err
		}

		return []domain.Clothes{}, http.StatusNotFound, err
	}

	return clothes, 0, nil
}
