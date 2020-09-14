package domain

import (
	"github.com/ZooArk/src/types"
	"github.com/gin-gonic/gin"
	"time"
)

// Category struct

type Category struct {
	Base
	Date *time.Time `json:"date"`
	Name string     `gorm:"type:varchar(30);not null" json:"name" binding:"required"`
}

type CategoryUsecase interface {
	Get(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// CategoryRepository is category interface for repository
type CategoryRepository interface {
	Add(category *Category) error
	Get(cateringID, clientID, date string) ([]Category, int, error)
	GetByKey(id, value, cateringID string) (Category, error)
	Delete(path types.PathCategory) (int, error)
	Update(path types.PathCategory, category *Category) (int, error)
}
