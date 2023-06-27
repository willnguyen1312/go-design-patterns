package patterns

import (
	"regexp"
	"strings"
)

// Interpreter interface
type Interpreter interface {
	Interprete(string) bool
}

// MatchIntepreter struct
type MatchIntepreter struct {
	Name string
}

// Interprete method on MatchIntepreter struct
func (i *MatchIntepreter) Interprete(str string) bool {
	r, _ := regexp.Compile("(?i)" + strings.ToLower(i.Name))
	return r.MatchString(str)
}

// CheckName funcion create a function to interprete the language passed in
func CheckName(name string) func(string) bool {
	matchCheck := &MatchIntepreter{
		Name: "React",
	}
	return func(str string) bool {
		return matchCheck.Interprete(str)
	}
}
