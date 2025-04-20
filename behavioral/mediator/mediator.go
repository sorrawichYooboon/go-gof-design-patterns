package mediator

import "fmt"

/*
==============================
Mediator Design Pattern
==============================

Definition:
Defines an object that encapsulates how a set of objects interact. This pattern promotes loose coupling by preventing objects from referring to each other explicitly.

Use Case:
Used when a set of objects communicate in a complex way and you want to centralize the communication logic.

Examples:
- Chat room systems
- UI form components interaction
- Air traffic control system

Why use it?
Reduces communication complexity and dependencies between collaborating objects.
*/

type Mediator interface {
	SendMessage(sender Colleague, message string)
}

type Colleague interface {
	Receive(message string)
}

type ChatRoom struct {
	users []Colleague
}

func (c *ChatRoom) Register(user Colleague) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) SendMessage(sender Colleague, message string) {
	for _, u := range c.users {
		if u != sender {
			u.Receive(message)
		}
	}
}

type User struct {
	Name     string
	mediator Mediator
}

func NewUser(name string, m Mediator) *User {
	return &User{Name: name, mediator: m}
}

func (u *User) Send(message string) {
	fmt.Printf("%s sends: %s\n", u.Name, message)
	u.mediator.SendMessage(u, message)
}

func (u *User) Receive(message string) {
	fmt.Printf("%s received: %s\n", u.Name, message)
}

func ExecuteMediatorPattern() {
	chatRoom := &ChatRoom{}

	alice := NewUser("Alice", chatRoom)
	bob := NewUser("Bob", chatRoom)
	charlie := NewUser("Charlie", chatRoom)

	chatRoom.Register(alice)
	chatRoom.Register(bob)
	chatRoom.Register(charlie)

	alice.Send("Hello, everyone!")
}
