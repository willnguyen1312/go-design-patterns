package patterns

func ExampleCommandButton() {
	browser := &Browser{}

	enableCommand := &EnableCommand{
		browser: browser,
	}
	disableCommand := &DisableCommand{
		browser: browser,
	}

	onButton := &CommandButton{
		renderer: enableCommand,
	}
	onButton.press()

	offButton := &CommandButton{
		renderer: disableCommand,
	}
	offButton.press()

	// Output:
	// Browser enable
	// Browser disable
}
