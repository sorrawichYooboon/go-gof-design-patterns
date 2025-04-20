package builder

import "fmt"

/*
==============================
Builder Design Pattern
==============================

Definition:
Separates the construction of a complex object from its representation, so the same construction process can create different representations.

Use Case:
Used when an object needs to be constructed step-by-step and might have many optional parts.

Examples:
- Creating HTML or XML documents
- Building a car with customizable features (engine, wheels, etc)
- Game character builder (weapons, skills, appearance)

Why use it?
Helps manage object creation complexity, especially when constructors are overloaded or nested.
*/

type Product struct {
	PartA string
	PartB string
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	GetResult() Product
}

type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.PartA = "Engine"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.PartB = "Wheels"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct() Product {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	return d.builder.GetResult()
}

// Usage
func ExecuteBuilderPattern() {
	builder := &ConcreteBuilder{}
	director := NewDirector(builder)

	product := director.Construct()

	fmt.Println("Product built with:")
	fmt.Println("- PartA:", product.PartA)
	fmt.Println("- PartB:", product.PartB)
}
