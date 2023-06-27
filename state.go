package patterns

import (
	"fmt"
)

// State interface
type State interface {
	TurnOn()
	TurnOff()
}

// LightBall struct
type LightBall struct {
	On           State
	Off          State
	CurrentState State
	Battery      int
}

// NewLightBall function retun a new Lightball
func NewLightBall(initBattery int) *LightBall {
	l := &LightBall{
		Battery: initBattery,
	}

	onState := &OnState{
		LightBall: l,
	}
	offState := &OffState{
		LightBall: l,
	}

	l.On = onState
	l.Off = offState

	l.SetState(offState)

	return l
}

// SetState method on LightBall struct
func (l *LightBall) SetState(s State) {
	l.CurrentState = s
}

// TurnOff method on LightBall struct
func (l *LightBall) TurnOff() {
	l.CurrentState.TurnOff()
}

// TurnOn method on LightBall struct
func (l *LightBall) TurnOn() {
	l.CurrentState.TurnOn()
}

// OnState struct
type OnState struct {
	LightBall *LightBall
}

// TurnOn method on OnState struct
func (o *OnState) TurnOn() {
	if o.LightBall.Battery > 0 {
		fmt.Println("OnState - TurnOn")
		o.LightBall.Battery--
		o.LightBall.SetState(o.LightBall.On)
	}
}

// TurnOff method on OnState struct
func (o *OnState) TurnOff() {
	if o.LightBall.Battery > 0 {
		fmt.Println("OnState - TurnOff")
		o.LightBall.Battery--
	}
}

// OffState struct
type OffState struct {
	LightBall *LightBall
}

// TurnOn method on OffState struct
func (o *OffState) TurnOn() {
	if o.LightBall.Battery > 0 {
		fmt.Println("OffState - TurnOn")
		o.LightBall.Battery--
		o.LightBall.SetState(o.LightBall.On)
	}
}

// TurnOff method on OffState struct
func (o *OffState) TurnOff() {
	if o.LightBall.Battery > 0 {
		fmt.Println("OffState - TurnOff")
		o.LightBall.Battery--
	}
}
