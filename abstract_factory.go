package patterns

import (
	"fmt"
)

// Button interface
type Button interface {
	Paint()
}

// GUIFactory interface
type GUIFactory interface {
	CreateButton() Button
}

// ReactFactory struct
type ReactFactory struct{}

// CreateButton method on ReactFactory struct
func (f *ReactFactory) CreateButton() Button {
	return new(ReactButton)
}

// VueFactory struct
type VueFactory struct{}

// CreateButton method on VueFactory struct
func (f *VueFactory) CreateButton() Button {
	return new(VueButton)
}

// AngularFactory struct
type AngularFactory struct{}

// CreateButton method on AngularFactory struct
func (f *AngularFactory) CreateButton() Button {
	return new(AngularButton)
}

// ReactButton struct
type ReactButton struct{}

// Paint method on ReactButton struct
func (b *ReactButton) Paint() {
	fmt.Println("ReactButton")
}

// VueButton struct
type VueButton struct{}

// Paint method on VueButton struct
func (b *VueButton) Paint() {
	fmt.Println("VueButton")
}

// AngularButton struct
type AngularButton struct{}

// Paint method on AngularButton struct
func (b *AngularButton) Paint() {
	fmt.Println("AngularButton")
}

const (
	// React denote
	React = iota
	// Vue denote
	Vue
	// Angular denote
	Angular
)

// BuildFactory function return GUIFactory instance
func BuildFactory(appearance int) GUIFactory {
	switch appearance {
	case React:
		return new(ReactFactory)
	case Vue:
		return new(VueFactory)
	case Angular:
		return new(AngularFactory)
	}
	return nil
}
