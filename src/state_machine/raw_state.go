package state_machine

import (
	"encoding/json"
	"bytes"
)

type RawState struct {
	Name 	string
	Id 		string
}

func NewRawState(msg string) *RawState {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_state RawState
	err := json.Unmarshal([]byte(msg), &raw_state)
	
	if err != nil {
		return nil
	}
	
	return &raw_state
}

func (this *RawState) Cmp(other *RawState) (equ bool) {
	if this.Name != other.Name { 
		return false 
	} else if this.Id != other.Id {
		return false 
	} 
	
	return true
}

func (this *RawState) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + this.Id)
	
	return buffer.String()
}
