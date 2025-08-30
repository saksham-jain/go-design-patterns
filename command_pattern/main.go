package main

import (
	"go-design-patterns/command_pattern/tv_command/button"
	"go-design-patterns/command_pattern/tv_command/commands/off_command"
	"go-design-patterns/command_pattern/tv_command/commands/on_command"
	"go-design-patterns/command_pattern/tv_command/switche"
	"go-design-patterns/command_pattern/tv_command/tv"
)

func main() {
	button1 := button.NewButton("Button1")
	button2 := button.NewButton("Button2")

	switch1 := switche.NewSwitch("Switch1")
	switch2 := switche.NewSwitch("Switche2")

	tv1 := tv.NewTv("TV1")
	commandOn := on_command.NewOnCommand("OnCommand", tv1)
	commandOff := off_command.NewOffCommand("OffCommand", tv1)
	button1.AssignCommand(commandOn)
	button2.AssignCommand(commandOff)

	switch1.AssignCommand(commandOn)
	switch2.AssignCommand(commandOff)

	button1.Press()
	button2.Press()

	switch2.Press()
	switch1.Press()
	switch2.Press()
}
