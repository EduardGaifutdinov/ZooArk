package main

import (
	"fmt"
	"github.com/ZooArk/db/seeds/dev"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/types"
)

func main() {
	fmt.Println("=== CREATING TYPES ===")
	createTypes()
	fmt.Println("=== TYPES ARE CREATED")

	migrate()
	fmt.Println("=== ADD MIGRATIONS ===")

	dev.CreateAdmin()
	dev.CreateUsers()
	dev.CreateProducts()
	dev.CreateCategory()
}

func migrate() {
	config.DB.DropTableIfExists(
		&domain.Base{},
		&domain.User{},
		&domain.Product{},
		&domain.Category{},
		&domain.Seed{},
	)

	config.DB.AutoMigrate(
		&domain.Seed{},
		&domain.Base{},
		&domain.User{},
		&domain.Category{},
		&domain.Product{},
	)

}

func addDbConstraints() {

}

func createTypes() {
	userTypesQuery := fmt.Sprintf("CREATE TYPE user_roles AS ENUM ('%s', '%s')",
		types.UserRoleEnum.SuperAdmin,
		types.UserRoleEnum.User,
	)

	config.DB.Exec(userTypesQuery)
}
