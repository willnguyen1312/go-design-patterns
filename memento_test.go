package patterns

import "fmt"

func ExampleMemento() {
	CareTaker := &CareTaker{
		mementoArray: make([]*Memento, 0),
	}
	TextArea := &TextArea{
		Content: "A",
	}

	TextArea.SetState("React")
	CareTaker.AddMemento(TextArea.CreateMemento())
	TextArea.SetState("Vue")
	CareTaker.AddMemento(TextArea.CreateMemento())
	TextArea.SetState("Angular")
	CareTaker.AddMemento(TextArea.CreateMemento())

	fmt.Println(CareTaker.GetMenento(0))
	fmt.Println(CareTaker.GetMenento(1))
	fmt.Println(CareTaker.GetMenento(2))

	//Output:
	// React
	// Vue
	// Angular
}
