package patterns

import (
	"fmt"
)

// RequesterAPI interface
type RequesterAPI interface {
	Request(string) error
}

// ReactService struct
type ReactService struct{}

// Request method on ReactService
func (s *ReactService) Request(msg string) error {
	fmt.Printf("ReactService: %s\n", msg)
	return nil
}

// VueService structy
type VueService struct {
}

// Request method on VueService
func (s *VueService) Request(msg string) error {
	fmt.Printf("VueService: %s\n", msg)
	return nil
}

// AngularService structy
type AngularService struct {
}

// Request method on AngularService
func (s *AngularService) Request(msg string) error {
	fmt.Printf("AngularService: %s\n", msg)
	return nil
}

// RequesterAbstraction interface
type RequesterAbstraction interface {
	CallAPI() error
}

// HTTPRequest struct
type HTTPRequest struct {
	URL       string
	Requester RequesterAPI
}

// CallAPI method on HTTPRequest struct
func (c *HTTPRequest) CallAPI() error {
	c.Requester.Request(c.URL)
	return nil
}
