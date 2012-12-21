package state_machine

import (
	"bytes"
	"strconv"
)

type StateTransition struct {
	Name 		string
	Id 			int
	From 		State
	To 			State
	Timestamp 	[]uint64
	Cause 		map[string]string
}

func NewStateTransition(raw_state_transition *RawStateTransition, states []State) *StateTransition {
	if raw_state_transition == nil {
		return nil
	}
	
	var state_transition StateTransition
	
	// Name 	
	
	// Id
	 
	// From
	
	// To
	
	// Cause
	
	
	return &state_transition
}

func (this *StateTransition) Cmp(other *StateTransition) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	} else if this.Timestamp != other.Timestamp {
		return false
	} else if !this.From.Cmp(&other.From) {
		return false
	} else if !this.To.Cmp(&other.To) {
		return false
	}
	
	for key, _ := range this.Cause {
		if _, ok := other.Cause[key]; !ok {
			return false
		}
		
		if this.Cause[key] != other.Cause[key] {
			return false
		}
	}

	return true
}

func (this *StateTransition) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString(" Id: " + strconv.Itoa(this.Id))
	buffer.WriteString(" Timestamp: " + strconv.FormatUint(this.Timestamp, 10))
	buffer.WriteString(" From: " + this.From.String())
	buffer.WriteString(" To: " + this.To.String())
	
	for key, value := range this.Cause {
		cause_str := "  Cause: " + key + "->" + value
		buffer.WriteString(cause_str)
	}
	
	return buffer.String()
}
