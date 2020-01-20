package patterns

import "fmt"

// Finder interface
type Finder interface {
	Search(string)
}

// Box struct
type Box struct {
	Components []Finder
	Name       string
}

// Search method on Box struct
func (f *Box) Search(item string) {
	fmt.Printf("Serching recursively for %s in box %s\n", item, f.Name)
	for _, composite := range f.Components {
		composite.Search(item)
	}
}

// Add method on Box struct
func (f *Box) Add(c Finder) {
	f.Components = append(f.Components, c)
}

// File struct
type File struct {
	Name string
}

// Search method on File struct
func (f *File) Search(item string) {
	fmt.Printf("Searching for item %s in file %s\n", item, f.Name)
}
