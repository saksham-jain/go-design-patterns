package button

import (
	"fmt"
	"go-design-patterns/command_pattern/tv_command"
)

type RemoteButton struct {
	name    string
	command tv_command.Command
}

func NewButton(name string) tv_command.Invoker {
	return &RemoteButton{
		name: name,
	}
}

func (b *RemoteButton) AssignCommand(command tv_command.Command) {
	b.command = command
}

func (b *RemoteButton) Press() {
	fmt.Printf("pressed %s\n", b.name)
	b.command.Execute()
}
