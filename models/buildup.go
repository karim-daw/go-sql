package models

type IBuildup interface {
	calculateVolume() float64
	calculateArea() float64
	calculateWeight() float64
	rotate(angle float64)
	calculateThickness() float64
}
