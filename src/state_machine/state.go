package state_machine

import (

)


type State struct {
	Name string
	Id int
	Timestamps []Timestamp
	Transitions []StateTransistion
}

func(state State) Test() string {
	return ""
}

