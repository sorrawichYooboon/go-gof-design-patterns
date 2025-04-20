package decorator

import "fmt"

/*
=================================
Decorator Design Pattern
=================================

Definition:
The Decorator pattern adds new responsibilities to objects dynamically without altering their structure. It’s a flexible alternative to subclassing for extending functionality.

Use Case:
Use when you want to add behavior to an object at runtime without modifying its class.

Examples:
- I/O stream enhancements (buffered, compressed)
- UI styling (scrollbars, borders)
- Logging or metrics wrappers

Why use it?
Supports open/closed principle: open for extension, closed for modification.
*/

// Component interface
type Notifier interface {
	Send(message string)
}

// Concrete Component
type EmailNotifier struct{}

func (n *EmailNotifier) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// Base Decorator
type NotifierDecorator struct {
	wrapped Notifier
}

func (d *NotifierDecorator) Send(message string) {
	d.wrapped.Send(message)
}

// Concrete Decorators
type SMSDecorator struct {
	NotifierDecorator
}

func (s *SMSDecorator) Send(message string) {
	s.NotifierDecorator.Send(message)
	fmt.Println("Sending SMS:", message)
}

type SlackDecorator struct {
	NotifierDecorator
}

func (s *SlackDecorator) Send(message string) {
	s.NotifierDecorator.Send(message)
	fmt.Println("Sending Slack:", message)
}

// Usage
func ExecuteDecoratorPattern() {
	email := &EmailNotifier{}
	sms := &SMSDecorator{NotifierDecorator{wrapped: email}}
	slack := &SlackDecorator{NotifierDecorator{wrapped: sms}}

	slack.Send("Hello, World!")
}
