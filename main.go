package main

import (
	"fmt"

	// Creational Patterns
	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/abstractfactory"
	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/builder"
	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/prototype"
	"github.com/sorrawichYooboon/go-gof-design-patterns/creational/singleton"

	// Structural Patterns
	"github.com/sorrawichYooboon/go-gof-design-patterns/structural/adapter"
	"github.com/sorrawichYooboon/go-gof-design-patterns/structural/bridge"
	"github.com/sorrawichYooboon/go-gof-design-patterns/structural/composite"
	"github.com/sorrawichYooboon/go-gof-design-patterns/structural/decorator"
	"github.com/sorrawichYooboon/go-gof-design-patterns/structural/facade"
	"github.com/sorrawichYooboon/go-gof-design-patterns/structural/flyweight"
	"github.com/sorrawichYooboon/go-gof-design-patterns/structural/proxy"

	// Behavioral Patterns
	"github.com/sorrawichYooboon/go-gof-design-patterns/behavioral/chainofresponsibility"
	"github.com/sorrawichYooboon/go-gof-design-patterns/behavioral/command"
	"github.com/sorrawichYooboon/go-gof-design-patterns/behavioral/iterator"
	"github.com/sorrawichYooboon/go-gof-design-patterns/behavioral/mediator"
	"github.com/sorrawichYooboon/go-gof-design-patterns/behavioral/memento"
	"github.com/sorrawichYooboon/go-gof-design-patterns/behavioral/observer"
	"github.com/sorrawichYooboon/go-gof-design-patterns/behavioral/visitor"
)

func main() {
	fmt.Println("=== Creational Patterns ===")
	runCreationalPatterns()

	fmt.Println("\n=== Structural Patterns ===")
	runStructuralPatterns()

	fmt.Println("\n=== Behavioral Patterns ===")
	runBehavioralPatterns()
}

func runCreationalPatterns() {
	fmt.Println("\n--- Abstract Factory ---")
	abstractfactory.ExecuteAbstractFactoryPattern()

	fmt.Println("\n--- Builder ---")
	builder.ExecuteBuilderPattern()

	fmt.Println("\n--- Prototype ---")
	prototype.ExecutePrototypePattern()

	fmt.Println("\n--- Singleton ---")
	singleton.ExecuteSingletonPattern()
}

func runStructuralPatterns() {
	fmt.Println("\n--- Adapter ---")
	adapter.ExecuteAdapterPattern()

	fmt.Println("\n--- Bridge ---")
	bridge.ExecuteBridgePattern()

	fmt.Println("\n--- Composite ---")
	composite.ExecuteCompositePattern()

	fmt.Println("\n--- Decorator ---")
	decorator.ExecuteDecoratorPattern()

	fmt.Println("\n--- Facade ---")
	facade.ExecuteFacadePattern()

	fmt.Println("\n--- Flyweight ---")
	flyweight.ExecuteFlyweightPattern()

	fmt.Println("\n--- Proxy ---")
	proxy.ExecuteProxyPattern()
}

func runBehavioralPatterns() {
	fmt.Println("\n--- Chain of Responsibility ---")
	chainofresponsibility.ExecuteChainOfResponsibilityPattern()

	fmt.Println("\n--- Command ---")
	command.ExecuteCommandPattern()

	fmt.Println("\n--- Iterator ---")
	iterator.ExecuteIteratorPattern()

	fmt.Println("\n--- Mediator ---")
	mediator.ExecuteMediatorPattern()

	fmt.Println("\n--- Memento ---")
	memento.ExecuteMementoPattern()

	fmt.Println("\n--- Observer ---")
	observer.ExecuteObserverPattern()

	fmt.Println("\n--- Visitor ---")
	visitor.ExecuteVisitorPattern()
}
