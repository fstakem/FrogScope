package state_machine

import (
	"bytes"
	"strconv"
	"log"
)

type TransitionEvent struct {
	Id 				int
	StateMachineId 	int
	FromStateId 		int
	ToStateId 		int
	Timestamp 		uint64
	Cause 			map[string]string
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
	transition_event.Cause = make(map[string]string)
	for key, value := range raw_transition_event.Cause {
		str_val, ok := value.(string)
		if ok {
			transition_event.Cause[key] = str_val
		}
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

func (this *TransitionEvent) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Id: " + strconv.Itoa(this.Id))
	buffer.WriteString("  StateMachineId: " + strconv.Itoa(this.StateMachineId))
	buffer.WriteString("  FromStateId: " + strconv.Itoa(this.FromStateId))
	buffer.WriteString("  ToStateId: " + strconv.Itoa(this.ToStateId))
	buffer.WriteString("  Timestamp: " + strconv.FormatUint(this.Timestamp, 10))
	
	for key, value := range this.Cause {
		cause_str := "  Cause: " + key + "->" + value
		buffer.WriteString(cause_str)
	}
	
	return buffer.String()
}

