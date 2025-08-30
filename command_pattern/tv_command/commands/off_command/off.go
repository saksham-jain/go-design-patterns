package off_command

import (
	"fmt"
	"go-design-patterns/command_pattern/tv_command"
)

type Off struct {
	name     string
	receiver tv_command.Receiver
}

func NewOffCommand(name string, receiver tv_command.Receiver) tv_command.Command {
	return &Off{
		name:     name,
		receiver: receiver,
	}
}

func (o *Off) Execute() {
	fmt.Printf("received execute, name: %s\n", o.name)
	o.receiver.TurnOff()
}
