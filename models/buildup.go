package models

type Buildup interface {
	calculateVolume() float64
	calculateArea() float64
	calculateWeight() float64
	calculateThickness() float64
	calculateUValue() float64
}
