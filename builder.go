package patterns

// BuildProcess interface
type BuildProcess interface {
	SetColor() BuildProcess
	SetTag() BuildProcess
	GetUIProduct() UIProduct
}

// UIDirector struct
type UIDirector struct {
	builder BuildProcess
}

// Construct method on UIDirector struct
func (f *UIDirector) Construct() {
	f.builder.SetTag().SetColor()
}

// SetBuilder method on UIDirector struct
func (f *UIDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// UIProduct struct
type UIProduct struct {
	Color string
	Tag   string
}

// ReactBuilder struct
type ReactBuilder struct {
	v UIProduct
}

// SetColor method on BuildProcess struct
func (c *ReactBuilder) SetColor() BuildProcess {
	c.v.Color = "Blue"
	return c
}

// SetTag method on BuildProcess struct
func (c *ReactBuilder) SetTag() BuildProcess {
	c.v.Tag = "React"
	return c
}

// GetUIProduct method on UIProduct struct
func (c *ReactBuilder) GetUIProduct() UIProduct {
	return c.v
}

// VueBuilder struct
type VueBuilder struct {
	v UIProduct
}

// SetColor method on BuildProcess struct
func (b *VueBuilder) SetColor() BuildProcess {
	b.v.Color = "Green"
	return b
}

// SetTag method on BuildProcess struct
func (b *VueBuilder) SetTag() BuildProcess {
	b.v.Tag = "Vue"
	return b
}

// GetUIProduct method on UIProduct struct
func (b *VueBuilder) GetUIProduct() UIProduct {
	return b.v
}

// AngularBuilder struct
type AngularBuilder struct {
	v UIProduct
}

// SetColor method on BuildProcess struct
func (b *AngularBuilder) SetColor() BuildProcess {
	b.v.Color = "Red"
	return b
}

// SetTag method on BuildProcess struct
func (b *AngularBuilder) SetTag() BuildProcess {
	b.v.Tag = "Angular"
	return b
}

// GetUIProduct method on UIProduct struct
func (b *AngularBuilder) GetUIProduct() UIProduct {
	return b.v
}
