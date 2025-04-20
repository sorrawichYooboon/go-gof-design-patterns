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
)

func main() {
	fmt.Println("=== Creational Patterns ===")
	runCreationalPatterns()

	fmt.Println("\n=== Structural Patterns ===")
	runStructuralPatterns()
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
