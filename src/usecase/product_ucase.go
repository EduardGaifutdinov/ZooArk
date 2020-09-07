package usecase

import (
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Product struct
type Product struct{}

// NewProduct returns pointer to product struct
// with all methods
func NewProduct() *Product {
	return &Product{}
}

var productRepo = repository.NewProductRepo()

// Add adds product with provided ID
// @tags product
// @Produce json
// @Param payload body request.AddProduct false "product object"
// @Success 200 {object} domain.Product false "product object"
// @Failure 400 {object} types.Error "Error"
// @Router /products [post]
func (p Product) Add(c *gin.Context) {
	var product domain.Product

	if err := utils.RequestBinderBody(&product, c); err != nil {
		return
	}

	err := productRepo.Add(&product)

	if err != nil {
		utils.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusOK, product)
}
