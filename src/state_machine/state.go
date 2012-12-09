package state_machine

import (

)


type State struct {
	Name string
	Timestamps []Timestamp
	Transitions []StateTransistion
}

func(state State) Test() string {
	return ""
}

