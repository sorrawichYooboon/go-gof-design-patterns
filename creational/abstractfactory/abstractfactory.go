package abstractfactory

import "fmt"

/*
========================================
Abstract Factory Design Pattern
========================================

Definition:
Provides an interface for creating families of related or dependent objects without specifying their concrete classes.

Use Case:
When you want to create UI elements or other related components that vary by platform or environment.

Examples:
- UI toolkit (e.g., Windows vs MacOS buttons and checkboxes)
- Game objects for different themes/skins
- Database connector factories (e.g., MySQLFactory, PostgreSQLFactory)

Why use it?
Promotes consistency among products and allows scalability across multiple product variants.
*/

func GetUIFactory(os string) GUIFactory {
	if os == "mac" {
		return &MacFactory{}
	}
	return &WinFactory{}
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

type MacButton struct{}

func (b *MacButton) Paint() {
	fmt.Println("Rendering Mac Button")
}

type MacCheckbox struct{}

func (c *MacCheckbox) Paint() {
	fmt.Println("Rendering Mac Checkbox")
}

type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (f *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

type WinButton struct{}

func (b *WinButton) Paint() {
	fmt.Println("Rendering Windows Button")
}

type WinCheckbox struct{}

func (c *WinCheckbox) Paint() {
	fmt.Println("Rendering Windows Checkbox")
}
