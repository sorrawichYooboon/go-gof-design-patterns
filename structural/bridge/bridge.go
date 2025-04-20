package bridge

import "fmt"

/*
=================================
Bridge Design Pattern
=================================

Definition:
The Bridge pattern decouples an abstraction from its implementation, allowing them to vary independently.

Use Case:
Use when you want to separate a large class or closely related classes into two separate hierarchies—abstraction and implementation—that can be developed independently.

Examples:
- UI frameworks where you want to decouple widgets from the operating systems.
- Remote control and TV (one is abstraction, one is implementation).

Why use it?
Improves scalability and maintainability by separating concerns.
*/

// Implementation defines the low-level interface
type Implementation interface {
	Operation() string
}

// ConcreteImplementationA provides one implementation
type ConcreteImplementationA struct{}

func (impl *ConcreteImplementationA) Operation() string {
	return "ConcreteImplementationA: Operation"
}

// ConcreteImplementationB provides another implementation
type ConcreteImplementationB struct{}

func (impl *ConcreteImplementationB) Operation() string {
	return "ConcreteImplementationB: Operation"
}

// Abstraction defines the high-level interface
type Abstraction struct {
	impl Implementation
}

func (a *Abstraction) Operation() string {
	return "Abstraction: " + a.impl.Operation()
}

// ExtendedAbstraction can extend Abstraction further
type ExtendedAbstraction struct {
	impl Implementation
}

func (ea *ExtendedAbstraction) Operation() string {
	return "ExtendedAbstraction: " + ea.impl.Operation()
}

// Usage
func ExecuteBridgePattern() {
	a := &Abstraction{impl: &ConcreteImplementationA{}}
	fmt.Println(a.Operation())

	b := &ExtendedAbstraction{impl: &ConcreteImplementationB{}}
	fmt.Println(b.Operation())
}
