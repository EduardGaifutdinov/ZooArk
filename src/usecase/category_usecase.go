package usecase

import (
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/schemes/request"
	"github.com/ZooArk/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Category struct
type Category struct{}

// NewCategory return pointer to Category struct
// with all methods
func NewCategory() *Category {
	return &Category{}
}

var categoryRepo = repository.NewCategoryRepo()

// Add add product category in DB
// returns 200 if success and 4xx if request failed
// @Summary Returns error if exists and 200 if success
// @Produce json
// @Accept json
// @Tags categories
// @Param body body swagger.AddCategory false "Category Name"
// @Success 200 {object} Category false "category object"
// @Failure 400 {object} types.Error "Error"
// @Router /categories [post]
func (pc Category) Add(c *gin.Context) {
	var body request.AddCategory

	if err := utils.RequestBinderBody(&body, c); err != nil {
		return
	}

	category := domain.Category{
		Date: body.Date,
		Name: body.Name,
	}

	err := categoryRepo.Add(&category)

	if err != nil {
		utils.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, category)
}