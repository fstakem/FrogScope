package state_machine

import (
	//"json"
)

type StateMachine struct {
	Name string
	Id int
	States []State
	Transitions []StateTransition
}

func NewStateMachine(data *map[string]interface{}) *StateMachine {
	if len(*data) == 0 {
		return nil
	}
	
	var state_machine StateMachine
	
	// Parse and fill in
	
	return &state_machine
}



