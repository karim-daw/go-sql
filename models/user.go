package models

import (
	"errors"
	"go-sql/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Posts    []Post `json:"posts" gorm:"foreignKey:AuthorID"`
}

func (user *User) Validate() error {
	if user.Username == "" {
		return errors.New("required username is missing")
	}
	return nil
}

// save user into database
func (user *User) Save() (*User, error) {
	err := database.DB.Create(&user).Error
	if err != nil {
		return &User{}, err // verbose error check if suer exists
	}
	return user, nil
}

// find first 100 users
func FindUsers() (*[]User, error) {
	users := []User{}
	err := database.DB.Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.DB.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
