package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;" json:"username"`
	Email    string `gorm:"size:255;not null;" json:"email"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

func (user *User) Save() (*User, error) {
	err := DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User

	if err := DB.Where("username = ?", username).Find(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}
