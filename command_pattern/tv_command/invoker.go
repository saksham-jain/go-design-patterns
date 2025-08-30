package tv_command

type Invoker interface {
	AssignCommand(command Command)
	Press()
}
