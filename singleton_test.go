package patterns

import "testing"

func TestGetHTTPRequestService(t *testing.T) {
	firstRequest := GetHTTPRequestService()
	if firstRequest == nil {
		t.Error("A new connection must have been made")
	}
	expectedCounter := firstRequest

	firstRequest.MakeRequest()

	if currentCount := firstRequest.GetNumberOfRequests(); currentCount != 1 {
		t.Errorf("After calling http service for the first time, the count must be 1 but it is %d\n", currentCount)
	}

	secondRequest := GetHTTPRequestService()
	if secondRequest != expectedCounter {
		t.Error("Singleton instances must be different")
	}

	secondRequest.MakeRequest()
	if currentCount := secondRequest.GetNumberOfRequests(); currentCount != 2 {
		t.Errorf("After calling http service for the second time using different request instance, the current count must be 2 but was %d\n", currentCount)
	}
}
