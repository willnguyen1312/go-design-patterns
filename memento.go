package patterns

// TextArea struct
type TextArea struct {
	Content string
}

// CreateMemento method of TextArea struct
func (e *TextArea) CreateMemento() *Memento {
	return &Memento{Content: e.Content}
}

// RestoreState method of TextArea struct
func (e *TextArea) RestoreState(m *Memento) {
	e.Content = m.GetSavedState()
}

// SetState method of TextArea struct
func (e *TextArea) SetState(Content string) {
	e.Content = Content
}

// GetState method of TextArea struct
func (e *TextArea) GetState() string {
	return e.Content
}

// Memento struct
type Memento struct {
	Content string
}

func (m *Memento) String() string {
	return m.Content
}

// GetSavedState method of Memento struct
func (m *Memento) GetSavedState() string {
	return m.Content
}

// CareTaker struct
type CareTaker struct {
	mementoArray []*Memento
}

// AddMemento method of CareTaker struct
func (c *CareTaker) AddMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

// GetMenento method of CareTaker struct
func (c *CareTaker) GetMenento(index int) *Memento {
	return c.mementoArray[index]
}
