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

	product := usecase.NewProduct()

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
		allUsers := authRequired.Group("/")
		allUsers.Use(validator.ValidateRoles(
			types.UserRoleEnum.SuperAdmin,
			types.UserRoleEnum.User,
			))
			{
				// Products
				allUsers.POST("/products", product.Add)
			}
	}
	return r
}
