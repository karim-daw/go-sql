package models

import "gorm.io/gorm"

type OpaqueMaterial struct {
	gorm.Model
	Name     string
	Quantity int
	GWP      float64
	Density  float64
}

func (o *OpaqueMaterial) setName(name string) string {
	return o.Name
}

func (o *OpaqueMaterial) setQuantity() int {
	return o.Quantity
}
