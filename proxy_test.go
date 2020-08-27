package patterns

func ExampleRendererProxy() {
	renderProxy := &RendererProxy{
		renderer: &ReactRenderer{},
	}

	renderProxy.Render()

	//Ouput:
	// I'mchecking ya...
	// I am React
}
