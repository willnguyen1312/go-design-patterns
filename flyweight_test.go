package patterns

import (
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	libraryFactory := NewLibraryFactory()

	libraries := make([]*Library, 500000*3)
	for i := 0; i < 500000; i = i + 3 {
		libraries[i] = libraryFactory.GetLibrary(ReactFlyweight)
		libraries[i+1] = libraryFactory.GetLibrary(VueFlyweight)
		libraries[i+2] = libraryFactory.GetLibrary(AngularFlyweight)
	}

	if libraryFactory.GetNumberOfObjects() != 3 {
		t.Errorf("The number of objects created was not 3: %d\n", libraryFactory.GetNumberOfObjects())
	}
}
