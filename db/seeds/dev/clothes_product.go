package dev

import (
	"fmt"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/utils"
	"sync"
)

// CreteClothes creates seeds for clothes table
func CreateClothes() {
	seedExists := config.DB.
		Where("name = ?", "init addresses").
		First(&domain.Seed{}).Error
	if seedExists != nil {
		seed := domain.Seed{
			Name: "init clothes",
		}

		var clothesArray []domain.Clothes
		productResult, _ := repository.NewProductRepo().
		utils.JSONParse("/db/seeds/data/clothes.json", &clothesArray)


		fmt.Println("=== Clothes seeds created ===")
	} else {
		fmt.Printf("Seed `init clothes` already exists \n")
	}
}