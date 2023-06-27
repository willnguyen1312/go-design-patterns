package patterns

func ExampleBox() {
	file1 := &File{Name: "file1"}
	folder1 := &Box{Name: "folder1"}
	folder2 := &Box{Name: "folder2"}

	folder1.Add(file1)
	folder2.Add(folder1)

	folder2.Search("Aha")

	// Output:
	// Serching recursively for Aha in box folder2
	// Serching recursively for Aha in box folder1
	// Searching for item Aha in file file1
}
