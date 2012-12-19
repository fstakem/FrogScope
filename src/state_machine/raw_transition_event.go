package state_machine

import (
	"encoding/json"
)

type RawTransitionEvent struct {
	Id 				string
	StateMachineId 	string
	FromStateId 		string
	ToStateId 		string
	Cause 			map[string]interface{}
}

func NewRawTransitionEvent(msg string) *RawTransitionEvent {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_transition_event RawTransitionEvent
	err := json.Unmarshal([]byte(msg), &raw_transition_event)
	
	if err != nil {
		return nil
	}
	
	return &raw_transition_event
}

func (this *RawTransitionEvent) Cmp(other *RawTransitionEvent) (equ bool) {
	if this.Id != other.Id { 
		return false 
	} else if this.StateMachineId != other.StateMachineId {
		return false 
	} else if this.FromStateId != other.FromStateId {
		return false 
	} else if this.ToStateId != other.ToStateId {
		return false 
	}
	
	return true
}

