package patterns

import "fmt"

// ResponsibilityLibrary interface
type ResponsibilityLibrary interface {
	Render(*UI)
	SetNext(ResponsibilityLibrary)
}

// ReactResponsibility struct
type ReactResponsibility struct {
	next ResponsibilityLibrary
}

// Render method on ReactResponsibility struct
func (r *ReactResponsibility) Render(ui *UI) {
	if ui.reactDone {
		fmt.Println("React Done")
		r.next.Render(ui)
		return
	}
	fmt.Println("React rendering UI")
	ui.reactDone = true
	r.next.Render(ui)
}

// SetNext method on ReactResponsibility struct
func (r *ReactResponsibility) SetNext(next ResponsibilityLibrary) {
	r.next = next
}

// VueResponsibility struct
type VueResponsibility struct {
	next ResponsibilityLibrary
}

// Render method on VueResponsibility struct
func (r *VueResponsibility) Render(ui *UI) {
	if ui.vueDone {
		fmt.Println("Vue Done")
		r.next.Render(ui)
		return
	}
	fmt.Println("Vue rendering UI")
	ui.vueDone = true
	r.next.Render(ui)
}

// SetNext method on VueResponsibility struct
func (r *VueResponsibility) SetNext(next ResponsibilityLibrary) {
	r.next = next
}

// AngularResponsibility struct
type AngularResponsibility struct {
	next ResponsibilityLibrary
}

// Render method on AngularResponsibility struct
func (r *AngularResponsibility) Render(ui *UI) {
	if ui.angularDone {
		fmt.Println("Angular Done")
		r.next.Render(ui)
		return
	}
	fmt.Println("Angular rendering UI")
	ui.angularDone = true
	r.next.Render(ui)
}

// SetNext method on AngularResponsibility struct
func (r *AngularResponsibility) SetNext(next ResponsibilityLibrary) {
	r.next = next
}

// FinalsResponsibility struct
type FinalsResponsibility struct {
	next ResponsibilityLibrary
}

// Render method on FinalsResponsibility struct
func (r *FinalsResponsibility) Render(ui *UI) {
	fmt.Println("Done Done Done")
}

// SetNext method on FinalsResponsibility struct
func (r *FinalsResponsibility) SetNext(next ResponsibilityLibrary) {
	r.next = next
}

// UI struct
type UI struct {
	reactDone   bool
	vueDone     bool
	angularDone bool
}
