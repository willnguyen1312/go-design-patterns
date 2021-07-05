package patterns

import "fmt"

func ExampleIterator() {
	react := &Library{
		Name: "React",
	}
	vue := &Library{
		Name: "Vue",
	}
	angular := &Library{
		Name: "Angular",
	}

	LibraryCollection := &LibraryCollection{
		Libraries: []*Library{react, vue, angular},
	}

	iterator := LibraryCollection.CreateIterator()
	for iterator.HasNext() {
		Library := iterator.GetNext()
		fmt.Println(Library)
	}

	// Output:
	// React
	// Vue
	// Angular
}
