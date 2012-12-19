package state_machine

import (
	"fmt"
	"encoding/json"
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

func NewTransitionEvent(msg string) *TransitionEvent {
	if len(msg) == 0 {
		return nil
	}

	var transition_event TransitionEvent
	err := json.Unmarshal([]byte(msg), &transition_event)
	
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	
	return &transition_event
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

