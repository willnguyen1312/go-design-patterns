package patterns

import (
	"fmt"
)

// Requester interface
type Requester interface {
	Request(url string) string
}

const (
	// ReactElement denote
	ReactElement = iota
	// VueElement denote
	VueElement
	// AngularElement denote
	AngularElement
)

//GetRenderUI returns a pointer to a Requester interface
func GetRenderUI(m int) (Requester, error) {
	switch m {
	case ReactElement:
		return new(ReactUI), nil
	case VueElement:
		return new(VueUI), nil
	case AngularElement:
		return new(AngularUI), nil
	default:
		return nil, fmt.Errorf("Renderer %d not recognized", m)
	}
}

// ReactUI struct
type ReactUI struct{}

// VueUI struct
type VueUI struct{}

// AngularUI struct
type AngularUI struct{}

// Request method on ReactUI struct
func (ui *ReactUI) Request(url string) string {
	return fmt.Sprintf("Request URL %s by using React\n", url)
}

// Request method on VueUI struct
func (ui *VueUI) Request(url string) string {
	return fmt.Sprintf("Request URL %s by using Vue\n", url)
}

// Request method on AngularUI struct
func (ui *AngularUI) Request(url string) string {
	return fmt.Sprintf("Request URL %s by using Angular\n", url)
}
