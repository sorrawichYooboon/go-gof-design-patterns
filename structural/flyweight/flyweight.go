package flyweight

import "fmt"

/*
=================================
Flyweight Design Pattern
=================================

Definition:
The Flyweight pattern minimizes memory usage by sharing as much data as possible with similar objects. It separates intrinsic (shared) from extrinsic (unique) state.

Use Case:
Use when you need to create a large number of similar objects and want to reduce memory usage.

Examples:
- Game particles or units (trees, bullets)
- Text character formatting
- Icon/image reuse in UI

Why use it?
Reduces the number of objects in memory by reusing existing instances.
*/

type TreeType struct {
	Name  string
	Color string
}

func (t *TreeType) Draw(x, y int) {
	fmt.Printf("Drawing tree [%s, %s] at (%d,%d)\n", t.Name, t.Color, x, y)
}

type TreeFactory struct {
	types map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{types: make(map[string]*TreeType)}
}

func (f *TreeFactory) GetTreeType(name, color string) *TreeType {
	key := name + "-" + color
	if _, exists := f.types[key]; !exists {
		f.types[key] = &TreeType{Name: name, Color: color}
	}
	return f.types[key]
}

type Tree struct {
	X, Y int
	Type *TreeType
}

func (t *Tree) Draw() {
	t.Type.Draw(t.X, t.Y)
}

// Usage
func ExecuteFlyweightPattern() {
	factory := NewTreeFactory()
	tree1 := Tree{X: 1, Y: 1, Type: factory.GetTreeType("Oak", "Green")}
	tree2 := Tree{X: 2, Y: 3, Type: factory.GetTreeType("Oak", "Green")}
	tree3 := Tree{X: 4, Y: 5, Type: factory.GetTreeType("Pine", "DarkGreen")}

	tree1.Draw()
	tree2.Draw()
	tree3.Draw()
}
