package proxy

import "fmt"

/*
=================================
Proxy Design Pattern
=================================

Definition:
The Proxy pattern provides a placeholder or surrogate for another object to control access to it.

Use Case:
Use when you want to add access control, lazy initialization, logging, or caching to an object without changing its implementation.

Examples:
- Virtual proxies (e.g., loading heavy objects)
- Access control (e.g., authorization check)
- Remote proxies (e.g., RPC communication)

Why use it?
Encapsulates logic around access or enhancements while preserving the original object's interface.
*/

type Subject interface {
	Request()
}

// RealSubject is the actual object
type RealSubject struct{}

func (r *RealSubject) Request() {
	fmt.Println("RealSubject: Handling request")
}

// Proxy controls access to RealSubject
type Proxy struct {
	real *RealSubject
}

func (p *Proxy) Request() {
	if p.real == nil {
		fmt.Println("Proxy: Initializing real subject...")
		p.real = &RealSubject{}
	}
	fmt.Println("Proxy: Delegating to real subject")
	p.real.Request()
}

// Usage
func ExecuteProxyPattern() {
	proxy := &Proxy{}
	proxy.Request()
}
