package dev

import (
	"fmt"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/utils"
)

// CreateProduct creates seeds for products table
func CreateProducts() {
	categoryRepo := repository.NewCategoryRepo()

	seedExists := config.DB.
		Where("name = ?", "init products").
		First(&domain.Seed{}).Error
	if seedExists != nil {
		seed := domain.Seed{
			Name: "init products",
		}

		var productsArray []domain.Product
		categoryResult, _ := categoryRepo.GetByKey("name", "супы")
		utils.JSONParse("/db/seeds/data/products.json", &productsArray)

		for i := range productsArray {
			productsArray[i].CategoryID = categoryResult.ID
			config.DB.Create(&productsArray[i])
		}
		config.DB.Create(&seed)
		fmt.Println("=== Products seeds created ====")
	} else {
		fmt.Printf("Seed `init products` already exists \n")
	}
}
