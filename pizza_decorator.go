package patterns

import (
	"fmt"
)

// LibraryDecorator interface
type LibraryDecorator interface {
	Decorate() string
}

// ReactLibrary struct
type ReactLibrary struct {
	Library LibraryDecorator
}

// Add method on ReactLibrary struct
func (p *ReactLibrary) Add(library LibraryDecorator) {
	p.Library = library
}

// Decorate method on ReactLibrary struct
func (p *ReactLibrary) Decorate() string {
	if p.Library != nil {
		s := p.Library.Decorate()
		return fmt.Sprintf("%s %s", s, "React")
	}
	return "React"
}

// VueLibrary struct
type VueLibrary struct {
	Library LibraryDecorator
}

// Add method on ReactLibrary struct
func (p *VueLibrary) Add(library LibraryDecorator) {
	p.Library = library
}

// Decorate method on VueLibrary struct
func (p *VueLibrary) Decorate() string {
	if p.Library != nil {
		s := p.Library.Decorate()
		return fmt.Sprintf("%s %s", s, "Vue")
	}
	return "Vue"
}

// AngularLibrary struct
type AngularLibrary struct {
	Library LibraryDecorator
}

// Add method on ReactLibrary struct
func (p *AngularLibrary) Add(library LibraryDecorator) {
	p.Library = library
}

// Decorate method on AngularLibrary struct
func (p *AngularLibrary) Decorate() string {
	if p.Library != nil {
		s := p.Library.Decorate()
		return fmt.Sprintf("%s %s", s, "Angular")
	}
	return "Angular"
}
