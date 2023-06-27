package patterns

import (
	"fmt"
)

// PizzaCook interface
type PizzaCook interface {
	GetIngredient()
	Cook()
	Serve()
}

// Cooker struct
type Cooker struct {
	PizzaCook PizzaCook
}

func (c *Cooker) makePizza() {
	c.PizzaCook.GetIngredient()
	c.PizzaCook.Cook()
	c.PizzaCook.Serve()
}

// ReactCooker struct
type ReactCooker struct{}

// GetIngredient method of  ReactCooker struct
func (r *ReactCooker) GetIngredient() {
	fmt.Println("React GetIngredient")
}

// Cook method of  ReactCooker struct
func (r *ReactCooker) Cook() {
	fmt.Println("React Cook")
}

// Serve method of  ReactCooker struct
func (r *ReactCooker) Serve() {
	fmt.Println("React Serve")
}

// VueCooker struct
type VueCooker struct{}

// GetIngredient method of  VueCooker struct
func (r *VueCooker) GetIngredient() {
	fmt.Println("Vue GetIngredient")
}

// Cook method of  VueCooker struct
func (r *VueCooker) Cook() {
	fmt.Println("Vue Cook")
}

// Serve method of  VueCooker struct
func (r *VueCooker) Serve() {
	fmt.Println("Vue Serve")
}
