package patterns

func ExampleMediator() {
	RenderingManager := NewRenderingManger()
	ReactRenderer := &ReactRenderer{
		Mediator: RenderingManager,
	}
	VueRenderer := &VueRenderer{
		Mediator: RenderingManager,
	}
	ReactRenderer.RequestRender()
	VueRenderer.RequestRender()
	ReactRenderer.FinishRender()

	// Output:
	// ReactRenderer: Rendering
	// VueRenderer: Waiting
	// ReactRenderer: Done
	// VueRenderer: Rendering Allowed, Rendering
}
