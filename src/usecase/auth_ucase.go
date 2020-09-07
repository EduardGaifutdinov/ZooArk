package usecase

import (
	"errors"
	"github.com/ZooArk/src/delivery/middleware"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth struct
type Auth struct{}

// NewAuth return pointer to Auth struct
// with all methods
func NewAuth() *Auth {
	return &Auth{}
}

// IsAuthenticated check if user is authorized and
// if user exists
// @Summary Returns user info if authorized
// @Produce json
// @Accept json
// @Tags auth
// @Success 200 {object} response.UserResponse
// @Failure 401 {object} types.Error
// @Failure 404 {object} types.Error
// @Router /is-authenticated [get]
func (a Auth) IsAuthenticated(c *gin.Context) {
	userRepo := repository.NewUserRepo()
	claims, err := middleware.Passport().CheckIfTokenExpire(c)

	if err != nil {
		utils.CreateError(http.StatusUnauthorized, err, c)
		return
	}

	if int64(claims["exp"].(float64)) < middleware.Passport().TimeFunc().Unix() {
		_, _, _ = middleware.Passport().RefreshToken(c)
	}

	id := claims[middleware.IdentityKeyID]
	result, err := userRepo.GetByID(id.(string))

	if err != nil {
		utils.CreateError(http.StatusUnauthorized, errors.New("token is expired"), c)
		return
	}

	if result.Status == &types.StatusTypesEnum.Deleted {
		utils.CreateError(http.StatusForbidden, errors.New("user was deleted"), c)
		return
	}

	c.JSON(http.StatusOK, result)

}

// @Summary Returns info about user
// @Produce json
// @Accept json
// @Tags auth
// @Param body body request.LoginUserRequest false "User Credentials"
// @Success 200 {object} response.UserResponse
// @Failure 401 {object} types.Error "Error"
// @Router /login [post]
// nolint:deadcode, unused
func login() {}

// @Summary Removes cookie if set
// @Produce json
// @Accept json
// @Tags auth
// @Success 200 {object} types.Error "Success"
// @Failure 401 {object} types.Error "Error"
// @Router /logout [get]
// nolint:deadcode, unused
func logout() {}
