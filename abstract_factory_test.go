package patterns

func ExampleBuildFactory() {
	if factory := BuildFactory(React); factory != nil {
		reactButton := factory.CreateButton()
		reactButton.Paint()
	}

	if factory := BuildFactory(Vue); factory != nil {
		vueButton := factory.CreateButton()
		vueButton.Paint()
	}

	if factory := BuildFactory(Angular); factory != nil {
		angularButton := factory.CreateButton()
		angularButton.Paint()
	}

	// Output:
	// ReactButton
	// VueButton
	// AngularButton
}
