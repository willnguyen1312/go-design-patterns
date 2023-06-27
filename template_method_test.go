package patterns

func ExamplePizzaCook() {
	reactCooker := &Cooker{
		PizzaCook: &ReactCooker{},
	}
	reactCooker.makePizza()

	vueCooker := &Cooker{
		PizzaCook: &VueCooker{},
	}
	vueCooker.makePizza()
	// Output:
	// React GetIngredient
	// React Cook
	// React Serve
	// Vue GetIngredient
	// Vue Cook
	// Vue Serve
}
