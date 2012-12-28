package state_machine

import (
	"bytes"
	"strconv"
	"log"
	"encoding/json"
)

type StateTransition struct {
	Name 		string
	Id 			int
	From 		State
	To 			State
	Timestamps 	[]uint64
	Causes 		[]Cause
}

func NewStateTransition(raw_state_transition *RawStateTransition, states []State) *StateTransition {
	if raw_state_transition == nil {
		return nil
	}
	
	var state_transition StateTransition
	
	// Name 	
	state_transition.Name = raw_state_transition.Name
	
	// Id
	id, err := strconv.Atoi(raw_state_transition.Id)
	if err != nil {
		log.Println("RawStateTransition.Id not properly parsed.")
		return nil
	}
	state_transition.Id = id
	 
	// From
	from_id, err := strconv.Atoi(raw_state_transition.FromId)
	if err != nil {
		log.Println("RawStateTransition.FromId not properly parsed.")
		return nil
	}
	
	var found_state bool = false
	for _, state := range states {
		if state.Id == from_id {
			state_transition.From = state
			found_state = true
		}
	}
	
	if !found_state {
		log.Println("StateTransition.From not found.")
		return nil
	}
	
	// To
	to_id, err := strconv.Atoi(raw_state_transition.ToId)
	if err != nil {
		log.Println("RawStateTransition.ToId not properly parsed.")
		return nil
	}
	
	found_state = false
	for _, state := range states {
		if state.Id == to_id {
			state_transition.To = state
			found_state = true
		}
	}
	
	if !found_state {
		log.Println("StateTransition.To not found.")
		return nil
	}
	
	// Causes
	j_data, err := json.Marshal(raw_state_transition.Causes)
	
	if err != nil {
		log.Println("RawStateTransition.Causes not properly parsed.")
		return nil
	}
	
	var raw_causes []interface{}
	data_str := string(j_data)
	err = json.Unmarshal([]byte(data_str), &raw_causes)
	
	if err != nil {
		log.Println("RawStateTransition.Causes not properly parsed.")
		return nil
	}
	
	for i := range raw_causes {
		j_data, err := json.Marshal(raw_causes[i])
	
		if err != nil {
			log.Println("RawStateTransition.Causes not properly parsed.")
			return nil
		}
		
		data_str := string(j_data)
		var raw_cause RawCause
		err = json.Unmarshal([]byte(data_str), &raw_cause)
		
		if err != nil {
			log.Println("RawStateTransition.Causes not properly parsed.")
			return nil
		}
	
		cause := NewCause(&raw_cause)
		state_transition.Causes = append(state_transition.Causes, *cause)
	}
	
	return &state_transition
}

func (this *StateTransition) Cmp(other *StateTransition) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	} else if !this.From.Cmp(&other.From) {
		return false
	} else if !this.To.Cmp(&other.To) {
		return false
	}
	
	if len(this.Causes) != len(other.Causes) {
		return false
	}
	
	for i := range this.Causes {
		if !this.Causes[i].Cmp(&other.Causes[i]) {
			return false
		}
	}

	return true
}

func (this *StateTransition) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString(" Id: " + strconv.Itoa(this.Id))
	buffer.WriteString(" From: " + this.From.String())
	buffer.WriteString(" To: " + this.To.String())
	
	buffer.WriteString("  Causes-> ")
	for i := range this.Causes {
		buffer.WriteString(this.Causes[i].String() + "  ")
	}
	
	return buffer.String()
}
