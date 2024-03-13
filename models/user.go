package models

import (
	"auth/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	RoleID   uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (user *User) Save() (*User, error) {
	err := pkg.DB.Create(&user).Error

	if err != nil {
		return &User{}, err
	}
	return user, nil
}
func GetUserByUsername(username string) (User, error) {
	var user User
	err := pkg.DB.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func GetUserById(id uint) (User, error) {
	var user User
	err := pkg.DB.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUser(user *User, id int) error {
	return pkg.DB.Where("id = ?", id).First(user).Error
}

func GetUsers(users *[]User) error {
	return pkg.DB.Find(users).Error
}

func UpdateUser(user *User) {
	_ = pkg.DB.Save(user)
}
