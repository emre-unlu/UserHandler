package dtos

type UserDto struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
}
