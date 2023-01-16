package models

import (
	"fmt"
)

// ---------------------------------

type NewMaterial struct {
	ID   int64
	Name string
	GWP  float64
}

// ---------------------------------

type Material interface {
	setGWP(gwp float64) bool
	setDensity(dens float64) bool
	materialInfo() string
}

type Envelope interface {
	setUValue(uVal float64) bool
	setThickness(thickness int32) bool
}

type Glass interface {
	setTransperancy(trans float64) bool
}

// ---------------------------------

func (m *NewMaterial) setUValue(uVal float64) bool {
	fmt.Printf("Im creating a new material named %s with a u-value: %0.2f\n W/m2", m.Name, uVal)
	return true
}

func (m *NewMaterial) setDensity(den float64) bool {
	fmt.Printf("Material named %s \t density: %0.2f\n kg/m3", m.Name, den)
	return true
}

func (m *NewMaterial) setThickness(thickness int32) bool {
	fmt.Printf("Material named %s \t thickness: %d\n", m.Name, thickness)
	return true
}

func (m *NewMaterial) setGWP(gwp float64) bool {
	fmt.Printf("This material named %s with a gwp: %0.2f kgco2e \n", m.Name, gwp)
	return true
}

func (m *NewMaterial) materialInfo() string {
	displayInfo := fmt.Sprintf("ID: %f, Name: %s, GWP: %0.2f \n", m.ID, m.Name, m.GWP)
	//fmt.Printf("ID: %s, Name: %s, GWP: %0.2f\n", m.ID, m.Name, m.GWP)
	return displayInfo
}

func (m *NewMaterial) setTransperancy(trans float64) bool {
	fmt.Printf("This material named %s with a transperancy: %0.2f\n", m.Name, trans)
	return true
}

// ---------------------------------

func AddGWP(material Material, gwp float64) {
	fmt.Println("adding CarbonArtifact...")
	material.setGWP(gwp)
	material.materialInfo()
}

func AddDensity(material Material, dens float64) {
	fmt.Println("adding CarbonArtifact...")
	material.setDensity(dens)
}

func AddThickness(envelop Envelope, thickness int32) {
	fmt.Println("adding Thickness...")
	envelop.setThickness(thickness)
}

func AddUval(envelop Envelope, uVal float64) {
	fmt.Println("adding Material...")
	envelop.setUValue(uVal)
}

func AddTransperancy(glass Glass, trans float64) {
	fmt.Println("adding Transperacy")
	glass.setTransperancy(trans)
}

// ---------------------------------

func main() {
	brick := &NewMaterial{
		ID:   1,
		Name: "Danish klinker brick type 01",
	}
	AddGWP(brick, 0.07)
	AddDensity(brick, 0.5)
	AddUval(brick, 0.2)
	AddThickness(brick, 200)

	// glass := &NewMaterial{
	// 	ID:   2,
	// 	Name: "Schuco: double glazing",
	// }
	// AddMaterial(glass, 0.47)
	// AddEnvelope(glass, 0.2)
	// AddGlassMaterial(glass, 0.6)
}
