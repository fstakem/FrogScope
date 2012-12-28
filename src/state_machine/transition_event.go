package state_machine

import (
	"bytes"
	"strconv"
	"log"
	"encoding/json"
)

type TransitionEvent struct {
	Id 				int
	StateMachineId 	int
	FromStateId 		int
	ToStateId 		int
	Timestamp 		uint64
	Causes		 	[]Cause
}

func NewTransitionEvent(raw_transition_event *RawTransitionEvent) *TransitionEvent {
	if raw_transition_event == nil {
		return nil
	}

	var transition_event TransitionEvent
	
	// Id
	id, err := strconv.Atoi(raw_transition_event.Id)
	if err != nil {
		log.Println("RawTransitionEvent.Id not properly parsed.")
		return nil
	}
	transition_event.Id = id
	
	// StateMachineId
	state_machine_id, err := strconv.Atoi(raw_transition_event.StateMachineId)
	if err != nil {
		log.Println("RawTransitionEvent.StateMachineId not properly parsed.")
		return nil
	}
	transition_event.StateMachineId = state_machine_id
	
	// FromStateId
	from_state_id, err := strconv.Atoi(raw_transition_event.FromStateId)
	if err != nil {
		log.Println("RawTransitionEvent.FromStateId not properly parsed.")
		return nil
	}
	transition_event.FromStateId = from_state_id
	
	// ToStateId
	to_state_id, err := strconv.Atoi(raw_transition_event.ToStateId)
	if err != nil {
		log.Println("RawTransitionEvent.ToStateId not properly parsed.")
		return nil
	}
	transition_event.ToStateId = to_state_id
	
	// Timestamp
	timestamp, err := strconv.ParseUint(raw_transition_event.Timestamp, 10, 64)
	if err != nil {
		log.Println("RawTransitionEvent.Timestamp not properly parsed.")
		return nil
	}
	transition_event.Timestamp = timestamp
	
	// Cause
	j_data, err := json.Marshal(raw_transition_event.Causes)
	
	if err != nil {
		log.Println("RawTransitionEvent.Causes not properly parsed.")
		return nil
	}
	
	var raw_causes []interface{}
	data_str := string(j_data)
	err = json.Unmarshal([]byte(data_str), &raw_causes)
	
	if err != nil {
		log.Println("RawTransitionEvent.Causes not properly parsed.")
		return nil
	}
	
	for i := range raw_causes {
		j_data, err := json.Marshal(raw_causes[i])
	
		if err != nil {
			log.Println("RawTransitionEvent.Causes not properly parsed.")
			return nil
		}
		
		data_str := string(j_data)
		var raw_cause RawCause
		err = json.Unmarshal([]byte(data_str), &raw_cause)
		
		if err != nil {
			log.Println("RawTransitionEvent.Causes not properly parsed.")
			return nil
		}
	
		cause := NewCause(&raw_cause)
		transition_event.Causes = append(transition_event.Causes, *cause)
	}
	
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

func (this *TransitionEvent) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Id: " + strconv.Itoa(this.Id))
	buffer.WriteString("  StateMachineId: " + strconv.Itoa(this.StateMachineId))
	buffer.WriteString("  FromStateId: " + strconv.Itoa(this.FromStateId))
	buffer.WriteString("  ToStateId: " + strconv.Itoa(this.ToStateId))
	buffer.WriteString("  Timestamp: " + strconv.FormatUint(this.Timestamp, 10))
	
	buffer.WriteString("  Causes-> ")
	for i := range this.Causes {
		buffer.WriteString(this.Causes[i].String() + "  ")
	}
	
	return buffer.String()
}

