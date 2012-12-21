package state_machine

import (
	"bytes"
	"strconv"
)

type StateMachine struct {
	Name string
	Id int
	States []State
	Transitions []StateTransition
}

func NewStateMachine(raw_state_machine *RawStateMachine) *StateMachine {
	if raw_state_machine == nil {
		return nil
	}
	
	var state_machine StateMachine
	
	// Parse and fill in
	// TODO
	//
	
	return &state_machine
}

func (this *StateMachine) Cmp(other *StateMachine) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	}
	
	if len(this.States) == len(other.States) {
		for i := range this.States {
			if !this.States[i].Cmp(&other.States[i]) {
				return false
			}
		}
	} else {
		return false
	}
	
	if len(this.Transitions) == len(other.Transitions) {
		for i := range this.Transitions {
			if !this.Transitions[i].Cmp(&other.Transitions[i]) {
				return false
			}
		}
	} else {
		return false
	}
	
	return true
}

func (this *StateMachine) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString(" Id: " + strconv.Itoa(this.Id))
	
	buffer.WriteString(" States: ")
	for i := range this.States {
		buffer.WriteString(this.States[i].String())
	}
	
	buffer.WriteString(" Transitions: ")
	for i := range this.Transitions {
		buffer.WriteString(this.Transitions[i].String())
	}
	
	return buffer.String()
}



