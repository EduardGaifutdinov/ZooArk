package repository

import (
	"github.com/ZooArk/src/config"
	"github.com/ZooArk/src/domain"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// UserRepo struct
type UserRepo struct {}

// NewUserRepo return pointer to user repository
// With all methods
func NewUserRepo() *UserRepo {
	return &UserRepo{}
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