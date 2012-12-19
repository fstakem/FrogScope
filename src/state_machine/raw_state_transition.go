package state_machine

import (
	"encoding/json"
	"bytes"
)

type RawStateTransition struct {
	Name 	string
	Id 		string
	FromId 	string
	ToId 	string
	Cause 	string
}

func NewRawStateTransition(msg string) *RawStateTransition {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_state_transition RawStateTransition
	err := json.Unmarshal([]byte(msg), &raw_state_transition)
	
	if err != nil {
		return nil
	}
	
	return &raw_state_transition
}

func (this *RawStateTransition) Cmp(other *RawStateTransition) (equ bool) {
	if this.Name != other.Name { 
		return false 
	} else if this.Id != other.Id {
		return false 
	} else if this.FromId != other.FromId {
		return false 
	} else if this.ToId != other.ToId {
		return false 
	} else if this.Cause != other.Cause {
		return false 
	} 
	
	return true
}

func (this *RawStateTransition) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + this.Id)
	buffer.WriteString("  FromId: " + this.FromId)
	buffer.WriteString("  ToId: " + this.ToId)
	buffer.WriteString("  Cause: " + this.Cause)
	
	return buffer.String()
}

