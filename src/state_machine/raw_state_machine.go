package state_machine

import (
	"encoding/json"
	"bytes"
)

type RawStateMachine struct {
	Name 		string
	Id 			string
	States 		map[string]interface{}
	Transitions 	map[string]interface{}
}

func NewRawStateMachine(msg string) *RawStateMachine {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_state_machine RawStateMachine
	err := json.Unmarshal([]byte(msg), &raw_state_machine)
	
	if err != nil {
		return nil
	}
	
	return &raw_state_machine
}

func (this *RawStateMachine) Cmp(other *RawStateMachine) (equ bool) {
	if this.Name != other.Name { 
		return false 
	} else if this.Id != other.Id {
		return false 
	} 
	
	return true
}

func (this *RawStateMachine) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + this.Id)
	
	return buffer.String()
}