package state_machine

import (
	"bytes"
	"strconv"
	"log"
	"encoding/json"
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
	
	// Name
	state_machine.Name = raw_state_machine.Name
	
	// Id
	id, err := strconv.Atoi(raw_state_machine.Id)
	if err != nil {
		log.Println("RawStateMachine.Id not properly parsed.")
		return nil
	}
	state_machine.Id = id
	
	// States
	j_data, err := json.Marshal(raw_state_machine.States)
	
	if err != nil {
		log.Println("RawStateMachine.States not properly parsed.")
		return nil
	}
	
	var raw_states []interface{}
	data_str := string(j_data)
	err = json.Unmarshal([]byte(data_str), &raw_states)
	
	if err != nil {
		log.Println("RawStateMachine.States not properly parsed.")
		return nil
	}
	
	for i := range raw_states {
		j_data, err := json.Marshal(raw_states[i])
	
		if err != nil {
			log.Println("RawStateMachine.States not properly parsed.")
			return nil
		}
		
		data_str := string(j_data)
		var raw_state RawState
		err = json.Unmarshal([]byte(data_str), &raw_state)
		
		if err != nil {
			log.Println("RawStateMachine.States not properly parsed.")
			return nil
		}
	
		state := NewState(&raw_state)
		state_machine.States = append(state_machine.States, *state)
	}
	
	// Transitions
	j_data, err = json.Marshal(raw_state_machine.Transitions)
	
	if err != nil {
		log.Println("RawStateMachine.Transitions not properly parsed.")
		return nil
	}
	
	var raw_transitions []interface{}
	data_str = string(j_data)
	err = json.Unmarshal([]byte(data_str), &raw_transitions)
	
	if err != nil {
		log.Println("RawStateMachine.Transitions not properly parsed.")
		return nil
	}
	
	for i := range raw_transitions {
		j_data, err := json.Marshal(raw_transitions[i])
	
		if err != nil {
			log.Println("RawStateMachine.Transitions not properly parsed.")
			return nil
		}
		
		data_str := string(j_data)
		var raw_state_transition RawStateTransition
		err = json.Unmarshal([]byte(data_str), &raw_state_transition)
		
		if err != nil {
			log.Println("RawStateMachine.Transitions not properly parsed.")
			return nil
		}
	
		state_transition := NewStateTransition(&raw_state_transition, &state_machine.States)
		state_machine.Transitions = append(state_machine.Transitions, *state_transition)
	}
	
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
		buffer.WriteString(this.States[i].String() + " ")
	}
	
	buffer.WriteString(" Transitions: ")
	for i := range this.Transitions {
		buffer.WriteString(this.Transitions[i].String() + " ")
	}
	
	return buffer.String()
}



