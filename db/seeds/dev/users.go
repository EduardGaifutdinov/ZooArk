package dev

import (
	"fmt"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/utils"
)

// CreateUsers will populate users table with random users
func CreateUsers() {
	seedExist := config.DB.Where("name = ?", "init users").First(&domain.Seed{}).Error
	if seedExist != nil {
		seed := domain.Seed{
			Name: "init users",
		}

		hashedPassword := utils.HashString("password12!")

		var userArray []domain.User
		utils.JSONParse("/db/seeds/data/users.json", &userArray)

		for i := range userArray {
			if i < 3 {
				userArray[i].Password = hashedPassword
				userArray[i].Status = &types.StatusTypesEnum.Active
				config.DB.Create(&userArray[i])
			} else {
				userArray[i].Password = hashedPassword
				userArray[i].Status = &types.StatusTypesEnum.Active
				config.DB.Create(&userArray[i])
			}
		}
		config.DB.Create(&seed)

		fmt.Println("=== User seeds created ===")
	} else {
		fmt.Printf("Seed `init users` already exists \n")
	}
}