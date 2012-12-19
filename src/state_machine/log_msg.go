package state_machine

import (
	"bytes"
	"strconv"
)

// Message enum type
type MsgType int

const (
	STATE_MACHINE = iota
	STATE_TRANSITION_EVENT
)

type LogMsg struct {
	Protocol 	string 				
	Version 		int					
	SeqNum 		int					
	Scope 		string				
	Timestamp 	uint64				
	Transmitter string				
	Type 		string				
	Data 		map[string]interface{}	
}

func NewLogMsg(raw_log_msg *RawLogMsg) *LogMsg {
	if raw_log_msg == nil {
		return nil
	}
	
	var log_msg LogMsg
	// TODO
	
	return &log_msg
}

func (this *LogMsg) Cmp(other *LogMsg) (equ bool) {
	if this.Protocol != other.Protocol { 
		return false 
	} else if this.Version != other.Version {
		return false 
	} else if this.SeqNum != other.SeqNum {
		return false
	} else if this.Scope != other.Scope {
		return false
	} else if this.Timestamp != other.Timestamp {
		return false
	} else if this.Type != other.Type {
		return false
	}
	
	return true
}

func (this *LogMsg) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Protocol: " + this.Protocol)
	buffer.WriteString("  Version: " + strconv.Itoa(this.Version))
	buffer.WriteString("  SeqNum: " + strconv.Itoa(this.SeqNum))
	buffer.WriteString("  Scope: " + this.Scope)
	buffer.WriteString("  Timestamp: " + strconv.FormatUint(this.Timestamp, 10))
	buffer.WriteString("  Transmitter: " + this.Transmitter)
	buffer.WriteString("  Type: " + this.Type)
	
	return buffer.String()
}
 
func ParseStateMachineMsg(data *map[string]interface{}) *StateMachine {
	return nil
}

func ParseStateTransitionMsg(data *map[string]interface{}) *StateTransition {
	/*
	for key, _ := range *data {
		if key == "Id" {
			fmt.Println("++++++++++++++++++", key)
		} else if key == "StateMachineId" {
			fmt.Println("++++++++++++++++++", key)
		} else if key == "FromStateId" {
			fmt.Println("++++++++++++++++++", key)
		} else if key == "ToStateId" {
			fmt.Println("++++++++++++++++++", key)
		} else if key == "Timestamp" {
			fmt.Println("++++++++++++++++++", key)
		} else if key == "Cause" {
			fmt.Println("++++++++++++++++++", key)
		}
	}
	*/
	return nil
}





