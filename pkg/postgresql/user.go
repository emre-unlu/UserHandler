package postgresql

import (
	"fmt"
	"github.com/emre-unlu/GinTest/internal/models"
	"gorm.io/gorm"
)

type PGUserRepository struct {
	DB *gorm.DB
}

func NewPGUserRepository(db *gorm.DB) *PGUserRepository {
	return &PGUserRepository{DB: db}
}

func (r *PGUserRepository) GetUserList(page uint, limit uint) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	offset := CalculateOffset(page, limit)

	// Count the total number of users
	if err := r.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		fmt.Sprintf("User count Error with error message : %w", err)
		return nil, 0, fmt.Errorf("internal error")
	}

	// Retrieve the list of users with pagination
	if err := r.DB.Limit(int(limit)).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve user list for page %d with limit %d: %w", page, limit, err)
	}

	return users, total, nil
}

func (r *PGUserRepository) GetUserById(id uint) (*models.User, error) {
	var user *models.User
	result := r.DB.First(&user, id)

	if result.Error != nil {
		return nil, fmt.Errorf("failed get the user: %w", result.Error)
	}

	return user, result.Error
}
func (r *PGUserRepository) CreateUser(user models.User) (*models.User, error) {
	result := r.DB.Create(&user)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to count total number of users: %w", result.Error)
	}

	return &user, result.Error
}

func (r *PGUserRepository) SuspendUserById(user *models.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to suspend the user with id %d: %w", user.ID, result.Error)
	}
	return nil
}

func (r *PGUserRepository) DeactivateUserById(user *models.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to deactivate the user with id %d: %w", user.ID, result.Error)
	}
	return nil
}

func (r *PGUserRepository) ActivateUserById(user *models.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		return fmt.Errorf("failed to save the user with id %d: %w", user.ID, result.Error)
	}
	return nil
}
func (r *PGUserRepository) UpdateUser(user *models.User) error {
	result := r.DB.Save(user)

	if result.Error != nil {
		return fmt.Errorf("failed to update the user with id %d: %w", user.ID, result.Error)
	}
	return nil
}

func (r *PGUserRepository) UpdatePassword(id uint, newPassword string) error {
	if err := r.DB.Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error; err != nil {
		return fmt.Errorf("failed to update password for user with ID %d: %w", id, err)
	}
	return nil
}
func (r *PGUserRepository) CheckUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ? AND status IN (?, ?)", email, models.StatusActive, models.StatusSuspended).
		First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to check user by email %s: %w", email, err)
	}
	return &user, nil
}
