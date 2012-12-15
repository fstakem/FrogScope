package state_machine

import (

)


type State struct {
	Name string
	Id int
	EnterTimestamp int64
	ExitTimestamp int64
	Transitions []StateTransistion
}

func(state State) Test() string {
	return ""
}

