package patterns

const (
	// ReactFlyweight denote
	ReactFlyweight = iota
	// VueFlyweight denote
	VueFlyweight
	// AngularFlyweight denote
	AngularFlyweight
)

// Library struct
type Library struct {
	ID           int
	Name         string
	DynamicStuff []byte
}

func getLibraryFactory(libraryCode int) Library {
	switch libraryCode {
	case ReactFlyweight:
		return Library{
			ID:   1,
			Name: "React",
		}
	case VueFlyweight:
		return Library{
			ID:   2,
			Name: "Vue",
		}
	case AngularFlyweight:
		return Library{
			ID:   3,
			Name: "Angular",
		}
	default:
		return Library{
			ID:   1,
			Name: "React",
		}
	}
}

// NewLibraryFactory to create LibraryFlyweightFactory
func NewLibraryFactory() LibraryFlyweightFactory {
	return LibraryFlyweightFactory{
		createdLibraries: make(map[int]*Library, 0),
	}
}

// LibraryFlyweightFactory struct
type LibraryFlyweightFactory struct {
	createdLibraries map[int]*Library
}

// GetLibrary method on LibraryFlyweightFactory struct
func (t *LibraryFlyweightFactory) GetLibrary(libraryCode int) *Library {
	if t.createdLibraries[libraryCode] != nil {
		return t.createdLibraries[libraryCode]
	}

	library := getLibraryFactory(libraryCode)
	t.createdLibraries[libraryCode] = &library

	return t.createdLibraries[libraryCode]
}

// GetNumberOfObjects method on LibraryFlyweightFactory struct
func (t *LibraryFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdLibraries)
}
