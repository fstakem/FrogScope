package state_machine

import (
	"encoding/json"
)

// Message enum type
type MsgType int

const (
	STATE_MACHINE = iota
	STATE_TRANSITION_EVENT
)

type LogMsg struct {
	Ver int
	SeqNum int
	Scope string
	Timestamp int64
	Type MsgType
	Data map[string]string
}


func NewLogMsg(msg string) *LogMsg {
	if len(msg) == 0 {
		return nil
	}
	
	log_msg := &LogMsg{}
	err := json.Unmarshal([]byte(msg), &log_msg)
	
	if err == nil {
		return nil
	}
	
	return log_msg
}
 

func ParseStateMachineMsg(msg string) *StateMachine {
	return nil
}

func ParseStateTransitionMsg(msg string) *StateTransition {
	return nil
}




