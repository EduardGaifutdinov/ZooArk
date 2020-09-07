package dev

import (
	"fmt"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/utils"
	"sync"
)

// CreateCategories creates seeds for categories table
func CreateCategory() {
	seedExists := config.DB.
		Where("name = ?", "init dish_categories").
		First(&domain.Seed{}).Error
	if seedExists != nil {
		seed := domain.Seed{
			Name: "init dish_categories",
		}

		var categoriesArray []domain.Category
		utils.JSONParse("/db/seeds/data/categories.json", &categoriesArray)

		var wg sync.WaitGroup
		wg.Add(len(categoriesArray))

		for i := range categoriesArray {
			go func(i int) {
				defer wg.Done()
				config.DB.Create(&categoriesArray[i])
			}(i)
		}

		wg.Wait()
		config.DB.Create(&seed)
		fmt.Println("=== Categories seeds created ===")
	} else {
		fmt.Printf("Seed `init dish_categories` already exists \n")
	}
}
