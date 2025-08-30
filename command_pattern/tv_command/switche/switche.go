package switche

import (
	"fmt"
	"go-design-patterns/command_pattern/tv_command"
)

type Switch struct {
	name    string
	command tv_command.Command
}

func NewSwitch(name string) tv_command.Invoker {
	return &Switch{
		name: name,
	}
}

func (s *Switch) AssignCommand(command tv_command.Command) {
	s.command = command
}

func (s *Switch) Press() {
	fmt.Printf("pressed %s\n", s.name)
	s.command.Execute()
}
