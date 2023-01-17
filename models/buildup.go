package models

import (
	"gorm.io/gorm"
)

type Buildup struct {
	gorm.Model
	Name      string
	Category  string
	Materials []Material `gorm:"many2many:buildup_materials;"`
}

type BuildupInterface interface {
	setName() string
	setCategory() bool
}
