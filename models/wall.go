package models

import (
	"gorm.io/gorm"
)

type Wall struct {
	gorm.Model
	Name        string
	Thickness   float64
	Materials   []Material
	Offset      float64
	Orientation float64
}

func (w *Wall) calculateVolume() float64 {
	var volume float64
	for _, material := range w.Materials {
		volume += material.getDensity() * material.getThickness()
	}
	return volume
}

func (w *Wall) calculateArea() float64 {
	var area float64
	for _, material := range w.Materials {
		area += material.getThickness()
	}
	return area
}

func (w *Wall) calculateWeight() float64 {
	return w.calculateVolume() * w.calculateArea()
}

func (w *Wall) calculateThickness() float64 {
	var thickness float64
	for _, material := range w.Materials {
		thickness += material.getThickness()
	}
	return thickness
}

func (w *Wall) addOpaqueMaterial(m *Material) {
	w.Materials = append(w.Materials, *m)
}

func (w *Wall) calculateUValue() float64 {
	var uValue float64
	for _, material := range w.Materials {
		uValue += material.getUValue()
	}
	return uValue
}
