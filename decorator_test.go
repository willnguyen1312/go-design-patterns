package patterns

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	react := &ReactLibrary{}
	vue := &VueLibrary{}
	angular := &AngularLibrary{}

	react.Add(vue)
	vue.Add(angular)

	got := react.Decorate()
	if want := "Angular Vue React"; got != want {
		t.Fatalf("Expected to get %s, but got %s", got, want)
	}
}
