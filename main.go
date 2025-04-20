package main

import (
	"fmt"

	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/abstractfactory"
	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/builder"
	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/prototype"
	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/singleton"
)

func main() {
	fmt.Println("== Creational Patterns ==")

	// Abstract Factory Pattern: Creates families of related objects
	uiFactory := abstractfactory.GetUIFactory("mac")
	button := uiFactory.CreateButton()
	checkbox := uiFactory.CreateCheckbox()
	button.Paint()
	checkbox.Paint()

	// Builder Pattern: Constructs complex objects step by step
	director := builder.NewDirector(&builder.ConcreteBuilder{})
	product := director.Construct()
	fmt.Println("Builder result:", product.PartA, "+", product.PartB)

	// Prototype Pattern: Clones an existing object
	original := &prototype.ConcretePrototype{Value: "original"}
	clone := original.Clone()
	fmt.Println("Prototype clone value:", clone.GetValue())

	// Singleton Pattern: Ensures a class has only one instance
	s := singleton.GetInstance()
	fmt.Println("Singleton message:", s.Message)
}
