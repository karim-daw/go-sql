package models

type IMaterial interface {
	getDensity() float64
	getThickness() float64
	getName() string
}

type Material struct {
	Name      string
	Color     string
	Density   float64
	Thickness float64
}

func (m *Material) getDensity() float64 {
	return m.Density
}

func (m *Material) getThickness() float64 {
	return m.Thickness
}

func (m *Material) getName() string {
	return m.Name
}
