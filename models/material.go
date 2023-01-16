package models

type IMaterial interface {
	getName() string
	getDensity() float64
	getThickness() float64
	displayMaterialInfo() string
}

type IOpaqueMaterial interface {
	getRValue() float64
}

type ITransperantMaterial interface {
	getTransperancy() float64
}

type Material struct {
	Name         string
	Color        string
	Density      float64
	Thickness    float64
	RValue       float64
	Transperancy float64
}

func (m *Material) getName() string {
	return m.Name
}

func (m *Material) getDensity() float64 {
	return m.Density
}

func (m *Material) getThickness() float64 {
	return m.Thickness
}

func (m *Material) displayMaterialInfo() string {
	return m.Name
}

func (m *Material) getRValue() float64 {
	return m.RValue
}
