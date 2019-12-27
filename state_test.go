package patterns

func ExampleState() {
	lightBall := NewLightBall(3)
	lightBall.TurnOn()
	lightBall.TurnOff()
	lightBall.TurnOn()

	// Useless - Battery is off
	lightBall.TurnOn()
	lightBall.TurnOn()
	lightBall.TurnOn()
	lightBall.TurnOn()

	//Output:
	// OffState - TurnOn
	// OnState - TurnOff
	// OnState - TurnOn
}
