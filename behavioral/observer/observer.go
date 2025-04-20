package observer

import "fmt"

/*
==============================
Observer Design Pattern
==============================

Definition:
Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified automatically.

Use Case:
Used when an object needs to notify others without being tightly coupled.

Examples:
- Event listeners (GUI, pub/sub systems)
- Stock price updates to subscribers
- Social media followers

Why use it?
Promotes loose coupling and dynamic subscription to changes.
*/

type Observer interface {
	Update(message string)
}

type Subject interface {
	Register(Observer)
	Deregister(Observer)
	Notify(message string)
}

type NewsPublisher struct {
	subscribers []Observer
}

func (n *NewsPublisher) Register(o Observer) {
	n.subscribers = append(n.subscribers, o)
}

func (n *NewsPublisher) Deregister(o Observer) {
	for i, obs := range n.subscribers {
		if obs == o {
			n.subscribers = append(n.subscribers[:i], n.subscribers[i+1:]...)
			break
		}
	}
}

func (n *NewsPublisher) Notify(message string) {
	for _, o := range n.subscribers {
		o.Update(message)
	}
}

type Subscriber struct {
	Name string
}

func (s *Subscriber) Update(message string) {
	fmt.Printf("%s received update: %s\n", s.Name, message)
}

func ExecuteObserverPattern() {
	news := &NewsPublisher{}
	alice := &Subscriber{Name: "Alice"}
	bob := &Subscriber{Name: "Bob"}

	news.Register(alice)
	news.Register(bob)

	news.Notify("New article on Design Patterns!")
}
