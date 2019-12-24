package patterns

import (
	"fmt"
)

// Renderer interface
type Renderer interface {
	Render()
}

// ReactRenderer struct
type ReactRenderer struct{}

// Render method on ReactRenderer struct
func (rr *ReactRenderer) Render() {
	fmt.Println("I am React")
}

// RendererProxy struct
type RendererProxy struct {
	renderer Renderer
}

// Render method on RendererProxy struct
func (rp *RendererProxy) Render() {
	fmt.Println("I'm checking ya...")
	rp.renderer.Render()
}
