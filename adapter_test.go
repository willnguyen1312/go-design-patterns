package patterns

func ExamplePrinterAdapter() {
	var printer Printer
	msg := "Hello Adapter Pattern!"

	printer = &PrinterAdapter{}
	printer.Print(msg)
	printer = &PrinterAdapter{Printer: &ReactLog{}}
	printer.Print(msg)

	// Output:
	// Hello Adapter Pattern!
	// Log from ReactLog: Hello Adapter Pattern!
}
