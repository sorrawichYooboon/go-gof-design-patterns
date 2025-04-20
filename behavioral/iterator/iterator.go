package iterator

import "fmt"

/*
==============================
Iterator Design Pattern
==============================

Definition:
Provides a way to access elements of a collection sequentially without exposing its underlying structure.

Use Case:
Used when you need to traverse a collection uniformly, regardless of how it's implemented.

Examples:
- Traversing lists, trees, or custom collections
- Database result sets

Why use it?
Promotes clean access to collection elements without exposing internal structure.
*/

type Iterator interface {
	HasNext() bool
	Next() string
}

type StringCollection struct {
	items []string
	index int
}

func NewStringCollection(items []string) *StringCollection {
	return &StringCollection{items: items}
}

func (s *StringCollection) HasNext() bool {
	return s.index < len(s.items)
}

func (s *StringCollection) Next() string {
	if s.HasNext() {
		val := s.items[s.index]
		s.index++
		return val
	}
	return ""
}

func ExecuteIteratorPattern() {
	collection := NewStringCollection([]string{"Go", "Design", "Patterns"})
	for collection.HasNext() {
		fmt.Println("Item:", collection.Next())
	}
}
