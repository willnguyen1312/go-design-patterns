package patterns

func ExampleResponsibilityLibrary() {
	ui := UI{}
	react := &ReactResponsibility{}
	vue := &VueResponsibility{}
	angular := &AngularResponsibility{}
	final := &FinalsResponsibility{}

	react.SetNext(vue)
	vue.SetNext(angular)
	angular.SetNext(final)

	react.Render(&ui)

	// Output:
	// React rendering UI
	// Vue rendering UI
	// Angular rendering UI
	// Done Done Done
}
