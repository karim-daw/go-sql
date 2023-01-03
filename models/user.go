package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"-"`
	Posts    []Post `json:"posts" gorm:"foreignKey:AuthorID"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("required name")
	}
	return nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// func (u *User) FindAllUsers(db *DB) (*[]User, error) {
func (u *User) FindUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}
