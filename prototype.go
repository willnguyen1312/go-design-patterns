package patterns

// Node interface
type Node interface {
	Clone() Node
}

// File struct
type File struct {
	Name string
}

// Clone method on File struct
func (f *File) Clone() Node {
	return &File{Name: f.Name}
}
