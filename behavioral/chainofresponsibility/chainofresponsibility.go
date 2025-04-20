package chainofresponsibility

import "fmt"

/*
===============================================
Chain of Responsibility Design Pattern
===============================================

Definition:
Allows a request to pass through a chain of handlers where each handler decides either to process the request or to pass it along the chain.

Use Case:
Used when multiple objects may handle a request, and the handler isn’t known beforehand.

Examples:
- Event handling in GUIs
- Logging levels (debug → info → warning → error)
- Middleware in web frameworks

Why use it?
Decouples sender and receiver of requests and promotes flexibility in assigning responsibilities.
*/

type Handler interface {
	SetNext(Handler)
	HandleRequest(string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(h Handler) {
	b.next = h
}

func (b *BaseHandler) PassToNext(request string) {
	if b.next != nil {
		b.next.HandleRequest(request)
	}
}

type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("Handler A handled the request.")
	} else {
		h.PassToNext(request)
	}
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("Handler B handled the request.")
	} else {
		h.PassToNext(request)
	}
}

func ExecuteChainOfResponsibilityPattern() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerA.SetNext(handlerB)

	fmt.Println("Sending request: A")
	handlerA.HandleRequest("A")

	fmt.Println("Sending request: B")
	handlerA.HandleRequest("B")

	fmt.Println("Sending request: C (unhandled)")
	handlerA.HandleRequest("C")
}
