package models

import "gorm.io/gorm"

type TransperancyMaterial struct {
	gorm.Model
	Name         string
	Quantity     int
	GWP          float64
	Density      float64
	Transperancy float64
}

func (o *TransperancyMaterial) getName() string {
	return o.Name
}

func (o *TransperancyMaterial) getQuantity() int {
	return o.Quantity
}
