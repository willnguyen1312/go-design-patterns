package patterns

func ExampleHTTPRequest() {
	httpRequest := &HTTPRequest{
		URL:       "URL",
		Requester: &ReactService{},
	}
	httpRequest.CallAPI()

	httpRequest = &HTTPRequest{
		URL:       "URL",
		Requester: &VueService{},
	}
	httpRequest.CallAPI()

	httpRequest = &HTTPRequest{
		URL:       "URL",
		Requester: &AngularService{},
	}
	httpRequest.CallAPI()

	// Output:
	// ReactService: URL
	// VueService: URL
	// AngularService: URL
}
