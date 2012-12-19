package state_machine

import (

)


type State struct {
	Name string
	Id int
	EnterTimestamp []int64
	ExitTimestamp []int64
}

func NewState(data *map[string]interface{}) *State {
	if len(*data) == 0 {
		return nil
	}
	
	var state State
	
	// Parse and fill in
	
	return &state
}

