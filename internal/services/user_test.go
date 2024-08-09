package services

import (
	"fmt"
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserById(id uint) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}
func (m *MockUserRepository) DeleteUserById(id uint) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(0)
}

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	mockUsers := []models.User{
		{ID: 1, Name: "John", Surname: "Doe", Email: "john@example.com", Phone: "1234567890", Birthdate: time.Now()},
		{ID: 2, Name: "Jane", Surname: "Doe", Email: "jane@example.com", Phone: "0987654321", Birthdate: time.Now()},
	}
	mockRepo.On("GetAllUsers").Return(mockUsers, nil)

	users, err := userService.GetAllUsers()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	userDTO := dtos.UserDto{
		Name:      "John",
		Surname:   "Doe",
		Email:     "john@example.com",
		Phone:     "1234567890",
		Birthdate: "2000-01-01",
	}

	birthdate, _ := time.Parse(internal.BirthdayFormat, userDTO.Birthdate)
	user := models.User{
		Name:      userDTO.Name,
		Surname:   userDTO.Surname,
		Email:     userDTO.Email,
		Phone:     userDTO.Phone,
		Birthdate: birthdate,
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("models.User")).Return(user, nil)

	createdUser, password, err := userService.CreateUser(userDTO)

	assert.Nil(t, err)
	assert.NotEmpty(t, password)
	assert.Equal(t, userDTO.Name, createdUser.Name)
	fmt.Printf("userDto : %+v , userModel : %+v\n", userDTO, createdUser)
	mockRepo.AssertExpectations(t)
}
