package usecase

import (
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/schemes/request"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

// Get returns list of categories or error
// @Summary Get list of categories
// @Tags categories
// @Produce json
// @Param date query string false "in format YYYY-MM-DDT00:00:00Z"
// @Success 200 {array} domain.Category "array of category readings"
// @Failure 400 {object} types.Error "Error"
// @Failure 404 {object} types.Error "Not Found"
// @Router /categories [get]
func (pc Category) Get(c *gin.Context) {
	var query types.DateQuery

	if err := utils.RequestBinderQuery(&query, c); err != nil {
		return
	}

	if query.Date == "" {
		query.Date = time.Now().Format(time.RFC3339)
	}

	categoriesResult, code, err := categoryRepo.Get(query.Date)

	if err != nil {
		utils.CreateError(code, err, c)
		return
	}

	c.JSON(http.StatusOK, categoriesResult)
}

// Delete soft delete of category reading
// @Summary Soft delete
// @Tags categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 204 "Successfully deleted"
// @Failure 404 {object} types.Error "Not Found"
// @Router /categories/{id} [delete]
func (pc Category) Delete(c *gin.Context) {
	var path types.PathCategory

	if err := utils.RequestBinderURI(&path, c); err != nil {
		return
	}

	if code, err := categoryRepo.Delete(path); err != nil {
		utils.CreateError(code, err, c)
		return
	}

	c.Status(http.StatusNoContent)
}

// Update updates category with new value provided in body
// @Summary Returns 204 if success and 4xx error if failed
// @Produce json
// @Accept json
// @Tags categories
// @Param id path string true "Category ID"
// @Param body body swagger.UpdateCategory false "new category name"
// @Success 204 "Successfully updated"
// @Failure 400 {object} types.Error "Error"
// @Failure 404 {object} types.Error "Not Found"
// @Router /categories/{id} [put]
func (pc Category) Update(c *gin.Context) {
	var path types.PathCategory
	var category domain.Category

	if err := utils.RequestBinderURI(&path, c); err != nil {
		return
	}

	if err := utils.RequestBinderBody(&category, c); err != nil {
		return
	}

	code, err := categoryRepo.Update(path, &category)

	if err != nil {
		utils.CreateError(code, err, c)
		return
	}

	c.Status(http.StatusNoContent)
}
