package patterns

func ExampleStrategy() {
	reactStrategy := &ReactStrategy{}
	vueStrategy := &VueStrategy{}
	angularStrategy := &AngularStrategy{}

	mobile := &Mobile{}
	str := "GUI"

	mobile.SetRenderStrategy(reactStrategy)
	mobile.Render(str)

	mobile.SetRenderStrategy(vueStrategy)
	mobile.Render(str)

	mobile.SetRenderStrategy(angularStrategy)
	mobile.Render(str)

	// Output:
	// React way: GUI
	// Vue way: GUI
	// Angular way: GUI
}
