package adapter

import "fmt"

/*
=================================
Adapter Design Pattern
=================================

Definition:
The Adapter pattern allows incompatible interfaces to work together by converting the interface of one class into an interface expected by the client.

Use Case:
Use this pattern when you want to use an existing class but its interface does not match the one you need.

Examples:
- Connecting legacy code with modern APIs.
- Adapting third-party libraries to your system.
- Power plug adapters (real-world analogy).

Why use it?
It promotes code reusability and compatibility without modifying existing source code.
*/

// Target is the expected interface
type Target interface {
	Request() string
}

// Adaptee is an existing class with a different interface
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee: Specific request"
}

// Adapter makes Adaptee compatible with Target
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

// Usage
func ExecuteAdapterPattern() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee}
	fmt.Println(adapter.Request())
}
