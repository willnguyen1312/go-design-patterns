package patterns

// Collection interface
type Collection interface {
	CreateIterator() Iterator
}

// LibraryCollection struct
type LibraryCollection struct {
	Libraries []*Library
}

// CreateIterator method on LibraryCollection struct
func (c *LibraryCollection) CreateIterator() Iterator {
	return &LibraryIterator{
		Libraries: c.Libraries,
	}
}

// Iterator interface
type Iterator interface {
	HasNext() bool
	GetNext() *Library
}

// LibraryIterator struct
type LibraryIterator struct {
	index     int
	Libraries []*Library
}

// HasNext method on LibraryIterator struct
func (i *LibraryIterator) HasNext() bool {
	return i.index < len(i.Libraries)
}

// GetNext method on LibraryIterator struct
func (i *LibraryIterator) GetNext() *Library {
	if i.HasNext() {
		Library := i.Libraries[i.index]
		i.index++
		return Library
	}
	return nil
}

// Library struct
type Library struct {
	Name string
}

func (l *Library) String() string {
	return l.Name
}
