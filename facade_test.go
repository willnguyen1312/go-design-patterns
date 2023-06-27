package patterns

func ExampleNewAllFacade() {
	allFacade := NewAllFacade()
	allFacade.reactTalk()
	allFacade.vueTalk()
	allFacade.angularTalk()

	// Output:
	// React
	// Vue
	// Angular
}
