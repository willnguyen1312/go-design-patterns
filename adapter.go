package patterns

import "fmt"

// Logger interface
type Logger interface {
	Log(string)
}

// ReactLog struct
type ReactLog struct{}

// Log method for ReactLog struct
func (l *ReactLog) Log(s string) {
	fmt.Printf("Log from ReactLog: %s", s)
}

// Printer interface
type Printer interface {
	Print(string)
}

// PrinterAdapter struct
type PrinterAdapter struct {
	Printer Logger
}

// Print method on PrinterAdapter struct
func (p *PrinterAdapter) Print(s string) {
	if p.Printer != nil {
		p.Printer.Log(s)
	} else {
		fmt.Println(s)
	}
}
