package memento

import "fmt"

/*
==============================
Memento Design Pattern
==============================

Definition:
Captures and externalizes an object’s internal state so it can be restored later, without violating encapsulation.

Use Case:
Used when an object’s state needs to be saved and restored, often for undo functionality.

Examples:
- Text editor undo functionality
- Game save/load system

Why use it?
Provides a way to restore an object’s state while preserving encapsulation.
*/

type Memento struct {
	State string
}

type Originator struct {
	State string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{State: o.State}
}

func (o *Originator) Restore(m *Memento) {
	o.State = m.State
}

type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index < 0 || index >= len(c.mementos) {
		return nil
	}
	return c.mementos[index]
}

func ExecuteMementoPattern() {
	originator := &Originator{State: "State1"}
	caretaker := &Caretaker{}

	caretaker.AddMemento(originator.CreateMemento())

	originator.State = "State2"
	caretaker.AddMemento(originator.CreateMemento())

	originator.State = "State3"
	fmt.Println("Current:", originator.State)

	originator.Restore(caretaker.GetMemento(0))
	fmt.Println("Restored to:", originator.State)
}
