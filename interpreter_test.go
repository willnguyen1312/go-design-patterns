package patterns

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	var tests = []struct {
		name string
		word string
		want bool
	}{
		{"React", "Sup react", true},
		{"Vue", "Yo bro Veu", false},
		{"Angular", "Anglar good?", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check := CheckName(tt.name)
			if got := check(tt.word); got != tt.want {
				t.Errorf("CheckName(%s) got %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
