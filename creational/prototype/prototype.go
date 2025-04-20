package prototype

/*
==============================
Prototype Design Pattern
==============================

Definition:
Creates new objects by copying an existing object (called the prototype).

Use Case:
Used when creating an object is costly (time/memory) and cloning is cheaper.

Examples:
- Game entities (e.g., cloning enemies/NPCs)
- Document templates
- GUI component duplication

Why use it?
Avoids expensive "new" object creation and simplifies instantiation for similar objects.
*/

type Prototype interface {
	Clone() Prototype
	GetValue() string
}

type ConcretePrototype struct {
	Value string
}

func (p *ConcretePrototype) Clone() Prototype {
	copy := *p
	return &copy
}

func (p *ConcretePrototype) GetValue() string {
	return p.Value
}
