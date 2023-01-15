package models

import "math"

type Wall struct {
	Name        string
	Thickness   float64
	Materials   []IMaterial
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

func (w *Wall) rotate(angle float64) {
	w.Orientation = math.Mod(w.Orientation+angle, 360.0)
}

func (w *Wall) calculateThickness() float64 {
	var thickness float64
	for _, material := range w.Materials {
		thickness += material.getThickness()
	}
	return thickness
}
