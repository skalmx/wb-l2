package pattern

import "fmt"

type Command interface {
	Execute()
}

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

type Receiver interface {
	StartEngine()
	TurnOffEngine()
}

type StartEngineCommand struct {
	r Receiver
}

func (c *StartEngineCommand) Execute() {
	c.r.StartEngine()
}

type TurnOffEngineCommand struct {
	r Receiver
}

func (c *TurnOffEngineCommand) Execute() {
	c.r.TurnOffEngine()
}

type Car struct {
}

func (c *Car) StartEngine() {
	fmt.Println("Your car is ready for ride")
}

func (c *Car) TurnOffEngine() {
	fmt.Println("The engine is off. Your trip is over")
}

// func main() {
// 	c := &Car{}

// 	s := &StartEngineCommand{
// 		r:c,
// 	}

// 	t := &TurnOffEngineCommand{
// 		r: c,
// 	}

// 	b := Button{
// 		command: s,
// 	}

// 	b1 := Button{
// 		command: t,
// 	}

// 	b.Press()
// 	b1.Press()
// }
