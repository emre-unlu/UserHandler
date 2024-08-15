package postgresql

import (
	"github.com/emre-unlu/GinTest/internal/models"
	"gorm.io/gorm"
)

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
func (r *PGUserRepository) DeleteUserById(id uint) (models.User, error) {
	result := r.DB.Delete(&models.User{}, id)
	return models.User{}, result.Error
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
	err := r.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
