package models

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name     string
	Quantity int
}

type MaterialInterface interface {
	setName() string
	setQuantity() int
}
