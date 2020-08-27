package main

import (
	"fmt"
	config "https:/src/config"
	"https:/src/domain"
	"https:/src/types"
)

func main()  {
	fmt.Println("=== CREATING TYPES ===")
	createTypes()
	fmt.Println("=== TYPES ARE CREATED")

	migrate()
	fmt.Println("=== ADD MIGRATIONS ===")
}

func migrate() {
	config.DB.DropTableIfExists(
		&domain.User{},
		&domain.Base{},
		)
	
	config.DB.AutoMigrate(
		&domain.User{},
		&domain.Base{},
		)
	
}

func addDbConstraints() {
	
}

func createTypes()  {
	userTypesQuery := fmt.Sprintf("CREATE TYPE user_roles AS ENUM ('%s', '%s', '%s', '%s')",
		types.UserRoleEnum.SuperAdmin,
		types.UserRoleEnum.User,
		)

	config.DB.Exec(userTypesQuery)
}
