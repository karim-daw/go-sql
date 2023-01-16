package models

type MaterialBuilder struct {
	Name         string
	Quantity     int
	GWP          float64
	Density      float64
	Transparency float64
	isOpaque     bool
}

func (b *MaterialBuilder) setName(name string) *MaterialBuilder {
	b.Name = name
	return b
}

func (b *MaterialBuilder) setQuantity(quantity int) *MaterialBuilder {
	b.Quantity = quantity
	return b
}

func (b *MaterialBuilder) setGWP(gwp float64) *MaterialBuilder {
	b.GWP = gwp
	return b
}

func (b *MaterialBuilder) setDensity(density float64) *MaterialBuilder {
	b.Density = density
	return b
}

func (b *MaterialBuilder) makeOpaque() *MaterialBuilder {
	b.isOpaque = true
	return b
}

func (b *MaterialBuilder) makeTransperant() *MaterialBuilder {
	b.isOpaque = false
	return b
}

// _____________________________

func (b *MaterialBuilder) Build() Material {
	if b.isOpaque {
		return &OpaqueMaterial{
			Name:     b.Name,
			Quantity: b.Quantity,
			GWP:      b.GWP,
			Density:  b.Density,
		}
	}
	return &TransparentMaterial{
		Name:         b.Name,
		Quantity:     b.Quantity,
		GWP:          b.GWP,
		Density:      b.Density,
		Transparency: b.Transparency,
	}
}
