package patterns

import "fmt"

// CommandButton struct
type CommandButton struct {
	renderer Renderer
}

func (b *CommandButton) press() {
	b.renderer.render()
}

// Renderer interface
type Renderer interface {
	render()
}

// EnableCommand struct
type EnableCommand struct {
	browser UI
}

func (c *EnableCommand) render() {
	c.browser.Enable()
}

// DisableCommand struct
type DisableCommand struct {
	browser UI
}

func (c *DisableCommand) render() {
	c.browser.Disable()
}

// UI interface
type UI interface {
	Enable()
	Disable()
}

// Browser struct
type Browser struct {
	enable bool
}

// Enable method on Browser struct
func (b *Browser) Enable() {
	b.enable = true
	fmt.Println("Browser enable")
}

// Disable method on Browser struct
func (b *Browser) Disable() {
	b.enable = false
	fmt.Println("Browser disable")
}
