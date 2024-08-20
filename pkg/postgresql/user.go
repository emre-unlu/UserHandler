package postgresql

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal/models"
	"gorm.io/gorm"
)

var ErrUserAlreadyDisactive = errors.New("user is already disactive")
var ErrUserAlreadyActive = errors.New("user is already active")
var ErrUserAlreadySuspended = errors.New("user is already suspended")
var ErrUserDeleted = errors.New("Deleted user cannot be reactivated")

type PGUserRepository struct {
	DB *gorm.DB
}

func NewPGUserRepository(db *gorm.DB) *PGUserRepository {
	return &PGUserRepository{DB: db}
}

func (r *PGUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	return users, result.Error
}
func (r *PGUserRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	return user, result.Error
}
func (r *PGUserRepository) CreateUser(user models.User) (models.User, error) {
	result := r.DB.Create(&user)
	return user, result.Error
}

func (r *PGUserRepository) SuspendUserById(user models.User) (models.User, error) {
	result := r.DB.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, result.Error
}

func (r *PGUserRepository) DeactivateUserById(user models.User) (models.User, error) {
	result := r.DB.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, result.Error
}

func (r *PGUserRepository) ActivateUserById(user models.User) (models.User, error) {
	result := r.DB.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, result.Error
}
func (r *PGUserRepository) UpdateUser(user models.User) (models.User, error) {
	result := r.DB.Save(&user)
	return user, result.Error
}

func (r *PGUserRepository) UpdatePassword(id uint, newPassword string) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
}
func (r *PGUserRepository) CheckUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ? AND status IN (?, ?)", email, models.StatusActive, models.StatusSuspended).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, nil // No user found
		}
		return models.User{}, err // Some other error occurred
	}
	return user, nil
}
