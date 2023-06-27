package patterns

import (
	"fmt"
)

// Renderer interface
type Renderer interface {
	RequestRender()
	FinishRender()
	AllowArrival()
}

// Mediator interface
type Mediator interface {
	canRender(Renderer) bool
	notifyReady()
}

// ReactRenderer struct
type ReactRenderer struct {
	Mediator Mediator
}

// RequestRender method on ReactRenderer struct
func (g *ReactRenderer) RequestRender() {
	if g.Mediator.canRender(g) {
		fmt.Println("ReactRenderer: Rendering")
	} else {
		fmt.Println("ReactRenderer: Waiting")
	}
}

// FinishRender method on ReactRenderer struct
func (g *ReactRenderer) FinishRender() {
	fmt.Println("ReactRenderer: Done")
	g.Mediator.notifyReady()
}

// AllowArrival method on ReactRenderer struct
func (g *ReactRenderer) AllowArrival() {
	fmt.Println("ReactRenderer: Rendering Allowed, Rendering")
}

// VueRenderer struct
type VueRenderer struct {
	Mediator Mediator
}

// RequestRender method on VueRenderer struct
func (g *VueRenderer) RequestRender() {
	if g.Mediator.canRender(g) {
		fmt.Println("VueRenderer: Rendering")
	} else {
		fmt.Println("VueRenderer: Waiting")
	}
}

// FinishRender method on VueRenderer struct
func (g *VueRenderer) FinishRender() {
	g.Mediator.notifyReady()
	fmt.Println("VueRenderer: Done")
}

// AllowArrival method on VueRenderer struct
func (g *VueRenderer) AllowArrival() {
	fmt.Println("VueRenderer: Rendering Allowed, Rendering")
}

// RenderingManager struct
type RenderingManager struct {
	isPlatformFree   bool
	renderersInQueue []Renderer
}

// NewRenderingManger create a new RenderingManager
func NewRenderingManger() *RenderingManager {
	return &RenderingManager{
		isPlatformFree: true,
	}
}

func (s *RenderingManager) canRender(t Renderer) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.renderersInQueue = append(s.renderersInQueue, t)
	return false
}

func (s *RenderingManager) notifyReady() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.renderersInQueue) > 0 {
		firstRendererInQueue := s.renderersInQueue[0]
		s.renderersInQueue = s.renderersInQueue[1:]
		firstRendererInQueue.AllowArrival()
	}
}
