package usecase

import (
	"errors"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/schemes/request"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

// User struct
type User struct{}

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

	if err := utils.RequestBinderBody(&body, c); err != nil {
		return
	}

	if err := copier.Copy(&user, &body); err != nil {
		utils.CreateError(http.StatusBadRequest, err, c)
		return
	}

	if ok := utils.IsEmailValid(user.Email); !ok {
		utils.CreateError(http.StatusBadRequest, errors.New("email is not valid"), c)
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
			utils.CreateError(http.StatusBadRequest, err, c)
			return
		}

		c.JSON(http.StatusOK, user)
	}

	for i := range existingUsers {
		if *existingUsers[i].Status != types.StatusTypesEnum.Deleted {
			utils.CreateError(http.StatusBadRequest, errors.New("user with that email already exist"), c)
			return
		}
	}

	user, userErr := userRepo.Add(user)

	if userErr != nil {
		utils.CreateError(http.StatusBadRequest, userErr, c)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Delete deletes user
// @Summary Returns error or 204 status code if success
// @Produce json
// @Accept json
// @Tags Users
// @Param id path string false "User ID"
// @Success 204 "Successfully deleted"
// @Failure 400 {object} types.Error "Error"
// @Failure 404 {object} types.Error "Error"
// @Router /users/{id} [delete]
func (u *User) Delete(c *gin.Context) {
	var path types.PathUser
	var user domain.User

	if err := utils.RequestBinderURI(&path, c); err != nil {
		return
	}

	parsedUserID, _ := uuid.FromString(path.Id)
	user.ID = parsedUserID
	user.Status = &types.StatusTypesEnum.Deleted
	deletedAt := time.Now().AddDate(0, 0, 21).Truncate(time.Hour * 24)
	user.DeletedAt = &deletedAt

	ctxUser, _ := c.Get("user")
	ctxUserRole := ctxUser.(domain.User).Role

	if user.ID == ctxUser.(domain.User).ID {
		utils.CreateError(http.StatusBadRequest, errors.New("can't delete yourself"), c)
		return
	}

	code, err := userRepo.Delete(user, ctxUserRole)

	if err != nil {
		utils.CreateError(code, err, c)
		return
	}

	c.Status(http.StatusNoContent)
}

// Update updates user
// @Summary Returns error or 200 status code if success
// @Produce json
// @Accept json
// @Tags Users
// @Param id path string false "User ID"
// @Param body body request.UserUpdate false "User"
// @Success 200 {object} response.UserResponse false "Catering user"
// @Failure 400 {object} types.Error "Error"
// @Failure 404 {object} types.Error "Error"
// @Router /users/{id} [put]
func (u *User) Update(c *gin.Context) {
	var user domain.User
	var path types.PathUser
	var body request.User

	if err := utils.RequestBinderURI(&path, c); err != nil {
		return
	}

	if err := utils.RequestBinderBody(&body, c); err != nil {
		return
	}

	if body.Email != "" {
		if ok := utils.IsEmailValid(body.Email); !ok {
			utils.CreateError(http.StatusBadRequest, errors.New("email is not valid"), c)
			return
		}
	}

	if err := copier.Copy(&user, &body); err != nil {
		utils.CreateError(http.StatusBadRequest, err, c)
		return
	}

	parsedUserID, _ := uuid.FromString(path.Id)
	user.ID = parsedUserID
	code, err := userRepo.Update(&user)

	if err != nil {
		utils.CreateError(code, err, c)
		return
	}

	updateUser, _ := userRepo.GetByID(path.Id)
	c.JSON(http.StatusOK, updateUser)
}
