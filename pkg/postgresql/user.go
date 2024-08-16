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

func (r *PGUserRepository) SuspendUserById(id uint) (models.User, error) {
	var user models.User
	userToSuspend := r.DB.First(&user, id)
	if userToSuspend.Error != nil {
		return user, userToSuspend.Error
	}

	if user.Status == models.StatusSuspended {
		return user, ErrUserAlreadySuspended
	}
	if user.Status == models.StatusInactive {
		return user, ErrUserDeleted
	}

	user.Status = models.StatusSuspended
	result := r.DB.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, result.Error
}

func (r *PGUserRepository) DeactivateUserById(id uint) (models.User, error) {
	var user models.User
	userToDelete := r.DB.First(&user, id)
	if userToDelete.Error != nil {
		return user, userToDelete.Error
	}

	if user.Status == models.StatusInactive {
		return user, ErrUserAlreadyDisactive
	}

	user.Status = models.StatusInactive
	result := r.DB.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, result.Error
}

func (r *PGUserRepository) ActivateUserById(id uint) (models.User, error) {
	var user models.User
	userToActivate := r.DB.First(&user, id)
	if userToActivate.Error != nil {
		return user, userToActivate.Error
	}

	if user.Status == models.StatusActive {
		return user, ErrUserAlreadyActive
	}
	if user.Status == models.StatusInactive {
		return user, ErrUserDeleted
	}

	user.Status = models.StatusActive
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
func (r *PGUserRepository) CheckUserByEmail(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.User{}).
		Where("email = ? AND status IN (?, ?)", email, models.StatusActive, models.StatusSuspended).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
