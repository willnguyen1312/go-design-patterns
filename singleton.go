package patterns

// HTTPRequest struct
type HTTPRequest struct {
	count int
}

var httpService *HTTPRequest

// GetHTTPRequestService function returns a HTTPRequest
func GetHTTPRequestService() *HTTPRequest {
	if httpService == nil {
		httpService = new(HTTPRequest)
	}

	return httpService
}

// MakeRequest method on HTTPRequest struct
func (r *HTTPRequest) MakeRequest() {
	r.count++
}

// GetNumberOfRequests method on HTTPRequest struct
func (r *HTTPRequest) GetNumberOfRequests() int {
	return r.count
}
