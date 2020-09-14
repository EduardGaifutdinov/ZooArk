package delivery

import (
	"github.com/ZooArk/src/delivery/middleware"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

// RedirectFunc wrapper for a Gin Redirect function
// which takes a route as a string and returns original Gin Redirect func
func RedirectFunc(route string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, route)
		c.Abort()
	}
}

// SetupRouter setting up gin router and config
func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := usecase.NewAuth()

	validator := middleware.NewValidator()

	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{os.Getenv("CLIENT_URL"), os.Getenv("CLIENT_MOBILE_URL")}

	clothes := usecase.NewClothes()
	user := usecase.NewUser()
	category := usecase.NewCategory()

	configCors.AllowCredentials = true
	r.Use(cors.New(configCors))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	dir, _ := os.Getwd()
	r.Use(static.Serve("/static/", static.LocalFile(dir+"src/static/images", true)))

	r.GET("/ZooArk/static/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/is-authenticated", auth.IsAuthenticated)
	r.POST("/login", middleware.Passport().LoginHandler)
	r.GET("/logout", middleware.Passport().LogoutHandler)

	authRequired := r.Group("/")
	authRequired.Use(middleware.Passport().MiddlewareFunc())
	{
		admin := authRequired.Group("/")
		admin.Use(validator.ValidateRoles(
			types.UserRoleEnum.SuperAdmin,
		))
		{
			// Products
			admin.POST("/products/clothes", clothes.Add)
			admin.DELETE("/products/clothes/:id", clothes.Delete)

			// Categories
			admin.POST("/categories", category.Add)
			admin.DELETE("/categories/:id", category.Delete)
			admin.PUT("/categories/:id", category.Update)
		}

		allUsers := authRequired.Group("/")
		allUsers.Use(validator.ValidateRoles(
			types.UserRoleEnum.User,
			types.UserRoleEnum.SuperAdmin,
		))
		{
			// Products
			allUsers.GET("/products/clothes", clothes.Get)
			// Users
			allUsers.POST("/users", user.Add)
			allUsers.DELETE("/users/:id", user.Delete)
			allUsers.PUT("/users/:id", user.Update)
			allUsers.GET("/categories", category.Get)
		}
	}
	return r
}
