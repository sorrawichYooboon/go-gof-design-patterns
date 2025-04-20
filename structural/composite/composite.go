package composite

import "fmt"

/*
=================================
Composite Design Pattern
=================================

Definition:
The Composite pattern composes objects into tree-like structures to represent part-whole hierarchies. It lets clients treat individual objects and compositions of objects uniformly.

Use Case:
Use when dealing with tree-like structures where individual and composite objects should be handled the same way.

Examples:
- File system (files and folders)
- UI component trees (buttons, panels, windows)
- HTML DOM

Why use it?
It simplifies client code by allowing consistent treatment of simple and complex elements.
*/

// Component defines an interface for all objects
type Component interface {
	Display(indent string)
}

// Leaf is a simple element (e.g., file)
type Leaf struct {
	Name string
}

func (l *Leaf) Display(indent string) {
	fmt.Println(indent + "- " + l.Name)
}

// Composite can hold other components (e.g., folder)
type Composite struct {
	Name     string
	Children []Component
}

func (c *Composite) Add(child Component) {
	c.Children = append(c.Children, child)
}

func (c *Composite) Display(indent string) {
	fmt.Println(indent + "+ " + c.Name)
	for _, child := range c.Children {
		child.Display(indent + "  ")
	}
}

// Usage
func ExecuteCompositePattern() {
	root := &Composite{Name: "Root"}
	file1 := &Leaf{Name: "File1"}
	file2 := &Leaf{Name: "File2"}

	subFolder := &Composite{Name: "SubFolder"}
	subFolder.Add(&Leaf{Name: "File3"})

	root.Add(file1)
	root.Add(file2)
	root.Add(subFolder)

	root.Display("")
}
