package tv

import (
	"fmt"
	"go-design-patterns/command_pattern/tv_command"
)

type TvState string

const (
	On  TvState = "ON"
	Off TvState = "OFF"
)

func (state TvState) String() {
	fmt.Printf("%s", state)
}

func (state TvState) IsOn() bool {
	return state == "ON"
}

func (state TvState) IsOff() bool {
	return state == "OFF"
}

type Tv struct {
	name  string
	state TvState
}

func NewTv(name string) tv_command.Receiver {
	return &Tv{
		name:  name,
		state: Off,
	}
}

func (tv *Tv) TurnOn() {
	currState := tv.state
	if currState.IsOn() {
		fmt.Printf("%s is already %s\n", tv.name, currState)
	} else {
		tv.state = On
		fmt.Printf("turned ON %s\n", tv.name)
	}
}

func (tv *Tv) TurnOff() {
	currState := tv.state
	if currState.IsOff() {
		fmt.Printf("%s is already %s\n", tv.name, currState)
	} else {
		tv.state = Off
		fmt.Printf("turned OFF %s\n", tv.name)
	}
}
