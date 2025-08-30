package on_command

import (
	"fmt"
	"go-design-patterns/command_pattern/tv_command"
)

type On struct {
	name     string
	receiver tv_command.Receiver
}

func NewOnCommand(name string, receiver tv_command.Receiver) tv_command.Command {
	return &On{
		name:     name,
		receiver: receiver,
	}
}

func (o *On) Execute() {
	fmt.Printf("received execute, name: %s\n", o.name)
	o.receiver.TurnOn()
}
