// package main

// import (
// 	"go-sql/controllers"
// 	"go-sql/database"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	database.ConnectDatabase()

// 	router.POST("/posts", controllers.CreatePost)
// 	router.POST("/register", controllers.RegisterUser)

// 	router.GET("/posts", controllers.FindPosts)
// 	router.GET("/posts/:id", controllers.FindPost) // here!
// 	router.PATCH("/posts/:id", controllers.UpdatePost)
// 	router.DELETE("/posts/:id", controllers.DeletePost)
// 	router.DELETE("/posts", controllers.DeleteAllPosts)
// 	router.Run("localhost:8080")
// }

// package models

// import (
// 	"fmt"
// )

// // ---------------------------------

// type NewMaterial struct {
// 	ID   int64
// 	Name string
// 	GWP  float64
// }

// // ---------------------------------

// type Material interface {
// 	materialInfo() string
// 	setGWP(gwp float64) bool
// 	setDensity(dens float64) bool
// }

// type Envelope interface {
// 	setUValue(uVal float64) bool
// 	setThickness(thickness int32) bool
// }

// type Glass interface {
// 	setTransperancy(trans float64) bool
// }

// // ---------------------------------

// func (m *NewMaterial) materialInfo(id int32, name string) string {
// 	displayInfo := fmt.Sprintf("\nID: %d, Name: %s\n", id, name)
// 	fmt.Printf(displayInfo)
// 	//fmt.Printf("ID: %s, Name: %s, GWP: %0.2f\n", m.ID, m.Name, m.GWP)
// 	return displayInfo
// }

// func (m *NewMaterial) setUValue(uVal float64) bool {
// 	fmt.Printf("U-value: %0.2f W/m2\n", uVal)
// 	return true
// }

// func (m *NewMaterial) setDensity(den float64) bool {
// 	fmt.Printf("Density: %0.2f kg/m3\n", den)
// 	return true
// }

// func (m *NewMaterial) setThickness(thickness int32) bool {
// 	fmt.Printf("Thickness: %d\n", thickness)
// 	return true
// }

// func (m *NewMaterial) setGWP(gwp float64) bool {
// 	fmt.Printf("GWP: %0.2f kgco2e\n", gwp)
// 	return true
// }

// func (m *NewMaterial) setTransperancy(trans float64) bool {
// 	fmt.Printf("Transperancy: %0.2f\n", trans)
// 	return true
// }

// ---------------------------------

// func AddGWP(material Material, gwp float64) {
// 	fmt.Println("adding CarbonArtifact...")
// 	material.setGWP(gwp)
// 	material.materialInfo()
// }

// func AddDensity(material Material, dens float64) {
// 	fmt.Println("adding Density...")
// 	material.setDensity(dens)
// }

// func AddUval(envelop Envelope, uVal float64) {
// 	fmt.Println("adding Material...")
// 	envelop.setUValue(uVal)
// }

// func AddThickness(envelop Envelope, thickness int32) {
// 	fmt.Println("adding Thickness...")
// 	envelop.setThickness(thickness)
// }

// func AddTransperancy(glass Glass, trans float64) {
// 	fmt.Println("adding Transperacy")
// 	glass.setTransperancy(trans)
// }

// ---------------------------------

// func main() {
// 	// 	ID:   1,
// 	// 	Name: "Danish klinker brick type 01",
// 	// }

// 	brick := &NewMaterial{}
// 	brick.materialInfo(0, "Danish klinker brick type 01")
// 	brick.setGWP(0.134)
// 	brick.setDensity(0.5)
// 	brick.setThickness(400)
// 	brick.setUValue(0.31)

// 	glass := &NewMaterial{}
// 	glass.materialInfo(0, "Schuco: double glazing")
// 	glass.setGWP(0.334)
// 	glass.setDensity(0.5)
// 	glass.setThickness(100)
// 	glass.setUValue(1.41)
// 	glass.setTransperancy(0.7)

// 	// AddGWP(brick, 0.07)
// 	// AddDensity(brick, 0.5)
// 	// AddUval(brick, 0.2)
// 	// AddThickness(brick, 400)

// 	// glass := &NewMaterial{
// 	// 	ID:   2,
// 	// 	Name: "Schuco: double glazing",
// 	// }
// 	// AddGWP(glass, 1.3)
// 	// AddDensity(glass, .98)
// 	// AddUval(glass, 1.2)
// 	// AddThickness(glass, 200)
// 	// AddTransperancy(glass, 0.6)
// }