package command

import "fmt"

/*
==============================
Command Design Pattern
==============================

Definition:
Encapsulates a request as an object, thereby allowing parameterization of clients with queues, requests, and operations.

Use Case:
Used when you need to queue or log requests, or support undoable operations.

Examples:
- Remote controls
- Transaction logs
- Macro recording

Why use it?
Decouples objects that produce commands from those that execute them.
*/

type Command interface {
	Execute()
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(c Command) {
	r.command = c
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func ExecuteCommandPattern() {
	light := &Light{}
	onCommand := &LightOnCommand{light}
	offCommand := &LightOffCommand{light}

	remote := &RemoteControl{}
	remote.SetCommand(onCommand)
	remote.PressButton()

	remote.SetCommand(offCommand)
	remote.PressButton()
}
