package usecase

import (
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/schemes/request"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"net/http"
)

// User struct
type User struct {}

// NewUser return pointer to user struct
// with all methods
func NewUser() *User {
	return &User{}
}

var userRepo = repository.NewUserRepo()

// Add creates user
// @Summary Returns error or 201 status code if success
// @Produce json
// @Accept json
// @Tags Users
// @Param body body request.User false "User"
// @Success 201 {object} response.User false "User"
// @Failure 400 {object} types.Error "Error"
// @Router /users [post]
func (u *User) Add(c *gin.Context) {
	var body request.User
	var user domain.User
	//var path types.PathID

	//if err := utils.RequestBinderURI(&path, c); err != nil {
	//	return
	//}

	if err := utils.RequestBinderBody(&body, c); err != nil {
		return
	}

	if err := copier.Copy(&user, &body); err != nil {
		utils.CreateError(http.StatusBadRequest, err.Error(), c)
			return
	}

	if ok := utils.IsEmailValid(user.Email); !ok {
		utils.CreateError(http.StatusBadRequest, "email is not valid", c)
			return
	}

	user.Role = types.UserRoleEnum.User
	user.Status = &types.StatusTypesEnum.Invited
	password := utils.GenerateString(10)
	user.Password = utils.HashString(password)

	existingUsers, err := userRepo.GetAllByKey("email", user.Email)

	if gorm.IsRecordNotFoundError(err) {
		_, err := userRepo.Add(*&user)

		if err != nil {
			utils.CreateError(http.StatusBadRequest, err.Error(), c)
			return
		}

		c.JSON(http.StatusOK, user)
	}

	for i := range existingUsers {
		if *existingUsers[i].Status != types.StatusTypesEnum.Deleted {
			utils.CreateError(http.StatusBadRequest, "user with that email already exist", c)
			return
		}
	}

	user, userErr := userRepo.Add(user)

	if userErr != nil {
		utils.CreateError(http.StatusBadRequest, userErr.Error(), c)
		return
	}

	c.JSON(http.StatusCreated, user)


}