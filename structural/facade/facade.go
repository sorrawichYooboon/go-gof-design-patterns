package facade

import "fmt"

/*
=================================
Facade Design Pattern
=================================

Definition:
The Facade pattern provides a simplified interface to a complex subsystem. It hides the complexity and exposes only necessary parts.

Use Case:
Use when you want to provide a unified API for a set of complex operations or components.

Examples:
- Simplified API over system libraries
- Startup/shutdown operations for applications
- Compiler interfaces (parsing, optimization, code generation)

Why use it?
Improves usability and decouples client code from internal system details.
*/

// Subsystems
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU started")
}

type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory loaded")
}

type Disk struct{}

func (d *Disk) Read() {
	fmt.Println("Disk read")
}

// Facade
type Computer struct {
	cpu    CPU
	memory Memory
	disk   Disk
}

func (c *Computer) Start() {
	c.cpu.Start()
	c.memory.Load()
	c.disk.Read()
}

// Usage
func ExecuteFacadePattern() {
	pc := &Computer{}
	pc.Start()
}
