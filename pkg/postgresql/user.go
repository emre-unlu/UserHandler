package postgresql

import (
	"fmt"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/pkg/postgresql/dto"
	"gorm.io/gorm"
)

type PGUserRepository struct {
	DB *gorm.DB
}

func NewPGUserRepository(db *gorm.DB) *PGUserRepository {
	return &PGUserRepository{DB: db}
}

func (r *PGUserRepository) GetUserList(page uint, limit uint, userFilterDto dto.UserFilter) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	offset := CalculateOffset(page, limit)

	query := r.DB.Model(&models.User{})

	filterConditions := dto.ConvertToFilterMap(userFilterDto)
	// Apply filters to the query
	for column, condition := range filterConditions {
		if condition.Value != nil {
			if condition.UseLike {
				query = query.Where(column+" ILIKE ?", "%"+*condition.Value+"%")
			} else {
				query = query.Where(column+" = ?", *condition.Value)
			}
		}
	}

	// Count the total number of users with filters applied
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Retrieve the list of users with pagination
	if err := query.Limit(int(limit)).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *PGUserRepository) GetUserById(id uint) (*models.User, error) {
	var user *models.User
	result := r.DB.First(&user, id)

	if result.Error != nil {

		fmt.Sprintf("Failed to retrive the user with code : %w", result.Error)
		return nil, fmt.Errorf("failed get the user")
	}

	return user, result.Error
}
func (r *PGUserRepository) CreateUser(user models.User) (*models.User, error) {
	result := r.DB.Create(&user)

	if result.Error != nil {
		fmt.Sprintf("Failed to create the user with id %d with error message: %w", user.ID, result.Error)
		return nil, fmt.Errorf("internal server error")
	}

	return &user, result.Error
}

func (r *PGUserRepository) SuspendUserById(user *models.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		fmt.Sprintf("Failed to suspend the user : %w", result.Error)
		return fmt.Errorf("failed to suspend the user with id %d", user.ID)
	}
	return nil
}

func (r *PGUserRepository) DeactivateUserById(user *models.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		fmt.Sprintf("Failed to deactivate the user : %w", result.Error)
		return fmt.Errorf("failed to deactivate the user with id %d", user.ID)
	}
	return nil
}

func (r *PGUserRepository) ActivateUserById(user *models.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		fmt.Sprintf("Failed to activate the user : %w", result.Error)
		return fmt.Errorf("failed to activate the user with id %d", user.ID)
	}
	return nil
}
func (r *PGUserRepository) UpdateUser(user *models.User) error {
	result := r.DB.Save(user)

	if result.Error != nil {
		fmt.Sprintf("Failed to update the user : %w", result.Error)
		return fmt.Errorf("failed to update the user with id %d: %w", user.ID, result.Error)
	}
	return nil
}

func (r *PGUserRepository) UpdatePassword(id uint, newPassword string) error {
	if err := r.DB.Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error; err != nil {
		fmt.Sprintf("Failed to update the password of  the user : %w", err)
		return fmt.Errorf("failed to update password for user with ID %d", id)
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
		fmt.Sprintf("Failed to checking the user : %w", err)
		return nil, fmt.Errorf("failed to check user by email %s", email)
	}
	return &user, nil
}
