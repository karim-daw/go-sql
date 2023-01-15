package models

import (
	"fmt"
	"math"
)

type Buildup interface {
	calculateVolume() float64
	calculateArea() float64
	calculateWeight() float64
	rotate(angle float64)
}

type Material interface {
	getDensity() float64
	getThickness() float64
}

type Wall struct {
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

func (w *Wall) rotate(angle float64) {
	w.Orientation = math.Mod(w.Orientation+angle, 360.0)
}

type Brick struct {
	Color     string
	Density   float64
	Thickness float64
}

func (b *Brick) getDensity() float64 {
	return b.Density
}

func (b *Brick) getThickness() float64 {
	return b.Thickness
}

func main() {
	wall := &Wall{
		Name:      "Wall",
		Thickness: 0.1,
		Materials: []Material{
			&Brick{Color: "Red", Density: 2000, Thickness: 0.05},
			&Brick{Color: "Gray", Density: 2500, Thickness: 0.05},
		},
		Offset:      0,
		Orientation: 0.0,
	}
	fmt.Println("Volume:", wall.calculateVolume())
	fmt.Println("Area:", wall.calculateArea())
	fmt.Println("Weight:", wall.calculateWeight())
	wall.rotate(90)
	fmt.Println("Orientation:", wall.Orientation)
}
