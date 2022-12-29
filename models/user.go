package models

import (
	"errors"
)

type User struct {
	ID    uint   `json:"user_id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Posts []Post `json:"posts" gorm:"foreignKey:AuthorID"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("required name")
	}
	return nil
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// func (u *User) FindAllUsers(db *DB) (*[]User, error) {
func (u *User) FindUsers() (*[]User, error) {
	var err error
	users := []User{}
	err = DB.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}
