package patterns

import "fmt"

// Platform interface
type Platform interface {
	GetType() string
	Accept(Visitor)
}

// Mobile struct
type Mobile struct{}

// Accept method on Mobile struct
func (s *Mobile) Accept(v Visitor) {
	v.VisitForMobile(s)
}

// GetType method on Mobile struct
func (s *Mobile) GetType() string {
	return "Mobile"
}

// Web struct
type Web struct{}

// Accept method on Web struct
func (c *Web) Accept(v Visitor) {
	v.VisitForWeb(c)
}

// GetType method on Web struct
func (c *Web) GetType() string {
	return "Web"
}

// Visitor interface
type Visitor interface {
	VisitForMobile(*Mobile)
	VisitForWeb(*Web)
}

// Renderer struct
type Renderer struct {
}

// VisitForMobile method on Renderer struct
func (a *Renderer) VisitForMobile(s *Mobile) {
	fmt.Println("Rendering on Mobile")
}

// VisitForWeb method on Renderer struct
func (a *Renderer) VisitForWeb(s *Web) {
	fmt.Println("Rendering on Web")
}
