package usecase

import (
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Clothes struct
type Clothes struct{}

// NewClothes returns pointer to product struct
// with all methods
func NewClothes() *Clothes {
	return &Clothes{}
}

var clothesRepo = repository.NewClothesRepo()

// Add adds product with provided ID
// @tags clothes
// @Produce json
// @Param payload body request.AddClothes false "clothes object"
// @Success 200 {object} domain.Clothes false "clothes object"
// @Failure 400 {object} types.Error "Error"
// @Router /products/clothes [post]
func (p Clothes) Add(c *gin.Context) {
	var product domain.Clothes

	category, _ := repository.NewCategoryRepo().GetByKey("name", "clothes")
	product.CategoryID = category.ID

	if err := utils.RequestBinderBody(&product, c); err != nil {
		return
	}

	err := clothesRepo.Add(&product)

	if err != nil {
		utils.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusCreated, product)
}

// Get return list of clothes
// @Summary Returns list of clothes
// @Tags clothes
// @Produce json
// @Success 200 {array} domain.Clothes "List of clothes"
// @Failure 400 {object} types.Error "Error"
// @Router /products/clothes [get]
func (p Clothes) Get(c *gin.Context) {
	clothes, code, err := clothesRepo.Get()

	if err != nil {
		utils.CreateError(code, err, c)
		return
	}

	if len(clothes) == 0 {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}

	c.JSON(http.StatusOK, clothes)
}
