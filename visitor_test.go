package patterns

func ExamplePlatform() {
	Mobile := &Mobile{}
	Web := &Web{}
	Renderer := &Renderer{}
	Mobile.Accept(Renderer)
	Web.Accept(Renderer)
	// Output:
	// Rendering on Mobile
	// Rendering on Web
}
