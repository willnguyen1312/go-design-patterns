package patterns

import "testing"

func TestClone(t *testing.T) {
	fileOriginal := &File{Name: "original"}

	fileClone := fileOriginal.Clone()

	fileCloneRaw, ok := fileClone.(*File)
	if !ok {
		t.Fatal("Type assertion for fileClone couldn't be done successfully")
	}

	fileOriginal.Name = "something awesome"
	if fileCloneRaw.Name == fileOriginal.Name {
		t.Error("fileOriginal cannot be equal to fileClone")
	}
}
