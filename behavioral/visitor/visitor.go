package visitor

import "fmt"

/*
==============================
Visitor Design Pattern
==============================

Definition:
Separates an algorithm from the object structure it operates on by allowing external visitors to perform operations.

Use Case:
Used when you need to perform many unrelated operations across different object structures without modifying those structures.

Examples:
- Operations on AST (Abstract Syntax Tree)
- Tax or billing calculation systems
- File system traversal

Why use it?
Promotes open/closed principle: new operations can be added without modifying the element classes.
*/

type Visitor interface {
	VisitBook(*Book)
	VisitDVD(*DVD)
}

type Item interface {
	Accept(Visitor)
}

type Book struct {
	Title string
	Price float64
}

func (b *Book) Accept(v Visitor) {
	v.VisitBook(b)
}

type DVD struct {
	Title string
	Price float64
}

func (d *DVD) Accept(v Visitor) {
	v.VisitDVD(d)
}

type DiscountVisitor struct{}

func (v *DiscountVisitor) VisitBook(b *Book) {
	fmt.Printf("Discounted Book: %s - $%.2f\n", b.Title, b.Price*0.9)
}

func (v *DiscountVisitor) VisitDVD(d *DVD) {
	fmt.Printf("Discounted DVD: %s - $%.2f\n", d.Title, d.Price*0.8)
}

func ExecuteVisitorPattern() {
	items := []Item{
		&Book{Title: "Go Programming", Price: 30.0},
		&DVD{Title: "Design Patterns", Price: 20.0},
	}

	visitor := &DiscountVisitor{}
	for _, item := range items {
		item.Accept(visitor)
	}
}
