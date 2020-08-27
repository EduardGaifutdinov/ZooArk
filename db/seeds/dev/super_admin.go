package dev

import (
	"fmt"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/utils"
)

const seedName string = "init admin"

// CreateAdmin init admin user
func CreateAdmin() {
	seedExists := config.DB.Where("name = ?", seedName).First(&domain.Seed{}).Error
	if seedExists != nil {
		seed := domain.Seed{
			Name: seedName,
		}

		superAdmin := domain.User{
			FirstName: "super",
			LastName:  "user",
			Email:     "admin@mail.ru",
			Password:  utils.HashString("password12!"),
			Status:    &types.StatusTypesEnum.Active,
			Role:      types.UserRoleEnum.SuperAdmin,
		}

		config.DB.Create(&superAdmin)
		config.DB.Create(&seed)
		fmt.Println("=== Admin seed created ===")
	} else {
		fmt.Printf("Seed `%s` already exists \n", seedName)
	}
}
