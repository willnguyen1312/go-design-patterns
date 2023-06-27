package patterns

import (
	"fmt"
)

// ReactFacade struct
type ReactFacade struct{}

// CallMe method on ReactFacade struct
func (r *ReactFacade) CallMe() {
	fmt.Println("React")
}

// VueFacade struct
type VueFacade struct{}

// CallMe method on ReactFacade struct
func (r *VueFacade) CallMe() {
	fmt.Println("Vue")
}

// AngularFacade struct
type AngularFacade struct{}

// CallMe method on ReactFacade struct
func (r *AngularFacade) CallMe() {
	fmt.Println("Angular")
}

// AllFacade struct
type AllFacade struct {
	React   *ReactFacade
	Vue     *VueFacade
	Angular *AngularFacade
}

func (af *AllFacade) reactTalk() {
	af.React.CallMe()
}

func (af *AllFacade) vueTalk() {
	af.Vue.CallMe()
}

func (af *AllFacade) angularTalk() {
	af.Angular.CallMe()
}

// NewAllFacade function to create a new AllFacade
func NewAllFacade() *AllFacade {
	return &AllFacade{
		React:   &ReactFacade{},
		Vue:     &VueFacade{},
		Angular: &AngularFacade{},
	}
}
