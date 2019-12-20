package patterns

import "testing"

func TestBuilderPattern(t *testing.T) {
	uiDirector := UIDirector{}

	reactBuilder := &ReactBuilder{}
	uiDirector.SetBuilder(reactBuilder)
	uiDirector.Construct()

	react := reactBuilder.GetUIProduct()

	if react.Color != "Blue" {
		t.Errorf("Color on a react ui must be 'Blue' and they were %s\n", react.Color)
	}

	if react.Tag != "React" {
		t.Errorf("Tag on a react ui must be 'React' and was %s\n", react.Tag)
	}

	VueBuilder := &VueBuilder{}

	uiDirector.SetBuilder(VueBuilder)
	uiDirector.Construct()

	vue := VueBuilder.GetUIProduct()

	if vue.Color != "Green" {
		t.Errorf("Color on a vue must be 'Green' and they were %s\n", vue.Color)
	}

	if vue.Tag != "Vue" {
		t.Errorf("Tag on a vue must be 'Vue' and was %s\n", vue.Tag)
	}

	angularBuilder := &AngularBuilder{}

	uiDirector.SetBuilder(angularBuilder)
	uiDirector.Construct()

	angular := angularBuilder.GetUIProduct()

	if angular.Color != "Red" {
		t.Errorf("Color on a angular must be 'Red' and they were %s\n", angular.Color)
	}

	if angular.Tag != "Angular" {
		t.Errorf("Tag on a angular must be 'Bus' and was %s\n", angular.Tag)
	}
}
