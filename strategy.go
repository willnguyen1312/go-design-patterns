package patterns

import (
	"fmt"
)

// Strategy interface
type Strategy interface {
	doWork(string)
}

// ReactStrategy struct
type ReactStrategy struct{}

func (s *ReactStrategy) doWork(message string) {
	fmt.Printf("React way: %s\n", message)
}

// VueStrategy struct
type VueStrategy struct{}

func (s *VueStrategy) doWork(message string) {
	fmt.Printf("Vue way: %s\n", message)
}

// AngularStrategy struct
type AngularStrategy struct{}

func (s *AngularStrategy) doWork(message string) {
	fmt.Printf("Angular way: %s\n", message)
}

// Platform interface
type Platform interface {
	Render(string)
}

// Mobile struct
type Mobile struct {
	renderStrategy Strategy
}

// Render method on Mobile struct
func (m *Mobile) Render(s string) {
	m.renderStrategy.doWork(s)
}

// SetRenderStrategy method on Mobile struct
func (m *Mobile) SetRenderStrategy(strategy Strategy) {
	m.renderStrategy = strategy
}
