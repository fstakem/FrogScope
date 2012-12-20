package state_machine

import (
	"bytes"
	"strconv"
)

type TransitionEvent struct {
	Id 				int
	StateMachineId 	int
	FromStateId 		int
	ToStateId 		int
	Timestamp 		int64
	Cause 			map[string]interface{}
}

func NewTransitionEvent(raw_transition_event *RawTransitionEvent) *TransitionEvent {
	if raw_transition_event == nil {
		return nil
	}

	var transition_event TransitionEvent
	// TODO 
	
	return &transition_event
}

func (this *TransitionEvent) Cmp(other *TransitionEvent) (equ bool) {
	if this.Id != other.Id { 
		return false 
	} else if this.StateMachineId != other.StateMachineId {
		return false 
	} else if this.FromStateId != other.FromStateId {
		return false
	} else if this.ToStateId != other.ToStateId {
		return false
	} else if this.Timestamp != other.Timestamp {
		return false
	} 
	
	return true
}

func (this *TransitionEvent) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Id: " + strconv.Itoa(this.Id))
	buffer.WriteString("  StateMachineId: " + strconv.Itoa(this.StateMachineId))
	buffer.WriteString("  FromStateId: " + strconv.Itoa(this.FromStateId))
	buffer.WriteString("  ToStateId: " + strconv.Itoa(this.ToStateId))
	buffer.WriteString("  Timestamp: " + strconv.FormatInt(this.Timestamp, 10))
	
	return buffer.String()
}

