package state_machine

import (

)

type StateTransition struct {
	Name string
	Id int
	timestamp int64
	From State
	To State
	Cause string
}

func NewStateTransition(data *map[string]interface{}) *StateTransition {
	if len(*data) == 0 {
		return nil
	}
	
	var state_transition StateTransition
	
	// Parse and fill in
	
	return &state_transition
}
