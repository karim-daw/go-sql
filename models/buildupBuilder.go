package models

import (
	"gorm.io/gorm"
)

type BuildupBuilder struct {
	gorm.Model
	Name      string
	Category  string
	Materials []Material
}

func (b *BuildupBuilder) setName(name string) *BuildupBuilder {
	b.Name = name
	return b
}

func (b *BuildupBuilder) setCategory(category string) *BuildupBuilder {
	b.Category = category
	return b
}

func (b *BuildupBuilder) Material(m Material) *BuildupBuilder {
	b.Materials = append(b.Materials, m)
	return b
}

func (b *BuildupBuilder) Build() *Buildup {
	return &Buildup{
		Name:      b.Name,
		Category:  b.Category,
		Materials: b.Materials,
	}
}
