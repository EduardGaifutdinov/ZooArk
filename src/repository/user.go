package repository

import (
	"errors"
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/ZooArk/src/types"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// UserRepo struct
type UserRepo struct{}

// NewUserRepo return pointer to user repository
// with all methods
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur UserRepo) GetAllByKey(key, value string) ([]domain.User, error) {
	var user []domain.User
	err := config.DB.
		Unscoped().
		Where(key+" = ?", value).
		Find(&user).Error

	return user, err
}

func (ur UserRepo) GetByID(id string) (domain.User, error) {
	var user domain.User
	err := config.DB.
		Table("users as u").
		Select("u.*").
		Where("u.id = ?", id).
		Scan(&user).
		Error

	config.DB.
		Table("users as u").
		Select("u.*").
		Where("u.id = ?", id).
		Scan(&user)

	return user, err
}

func (ur UserRepo) GeByKey(key, value string) (domain.User, error) {
	var user domain.User
	err := config.DB.
		Unscoped().
		Where(key+" = ?", value).
		First(&user).Error

	return user, err
}

func (ur UserRepo) UpdateStatus(userID uuid.UUID, status string) (int, error) {
	if err := config.DB.
		Model(&domain.User{}).
		Where("id = ?", userID).
		Update("status", status).
		Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return http.StatusNotFound, err
		}
		return http.StatusBadRequest, err
	}
	return 0, nil
}

func (ur UserRepo) Add(user domain.User) (domain.User, error) {

	if err := config.DB.
		Create(&user).
		Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur UserRepo) Delete(user domain.User, ctxUserRole string) (int, error) {
	var totalUsers int
	if ctxUserRole != types.UserRoleEnum.SuperAdmin {
		config.DB.
			Table("users as u").
			Where("u.status != ?", types.StatusTypesEnum.Deleted).
			Count(&totalUsers)

		if totalUsers == 1 {
			return http.StatusBadRequest, errors.New("can't delete last user")
		}
	}

	if userExist := config.DB.
		Table("users as u").
		Where("u.id = ?", user.ID).
		Update(&user).
		RowsAffected; userExist == 0 {
		return http.StatusBadRequest, errors.New("user not found")
	}

	return 0, nil
}

func (ur UserRepo) Update(user *domain.User) (int, error) {
	if userExist := config.DB.
		Where("id = ? AND email = ?", user.ID, user.Email).
		Find(&domain.User{}).
		RowsAffected; userExist == 0 {
		if emailExist := config.DB.
			Where("email = ?", user.Email).
			Find(&domain.User{}).
			RowsAffected; emailExist != 0 {
			return http.StatusBadRequest, errors.New("user with that email already exists")
		}
	}

	if err := config.DB.
		Unscoped().
		Model(&domain.User{}).
		Update(user).
		Error; err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return http.StatusNotFound, errors.New("user not found")
		}

		return http.StatusBadRequest, err
	}
	return 0, nil
}
