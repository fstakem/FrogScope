package state_machine

import (
	"bytes"
	"strconv"
	"log"
	"encoding/json"
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
	Data 		*map[string]interface{}	
}

func NewLogMsg(raw_log_msg *RawLogMsg) *LogMsg {
	if raw_log_msg == nil {
		return nil
	}
	
	var log_msg LogMsg
	
	// Protocol
	log_msg.Protocol = raw_log_msg.Protocol
	
	// Version
	version, err := strconv.Atoi(raw_log_msg.Version)
	if err != nil {
		log.Println("RawLogMsg.Version not properly parsed.")
		return nil
	}
	log_msg.Version = version
	
	// SeqNum
	seqNum, err := strconv.Atoi(raw_log_msg.SeqNum)
	if err != nil {
		log.Println("RawLogMsg.SeqNum not properly parsed.")
		return nil
	}
	log_msg.SeqNum = seqNum
	
	// Scope
	log_msg.Scope = raw_log_msg.Scope
	
	// Timestamp
	timestamp, err := strconv.ParseUint(raw_log_msg.Timestamp, 10, 64)
	if err != nil {
		log.Println("RawLogMsg.Timestamp not properly parsed.")
		return nil
	}
	log_msg.Timestamp = timestamp
	
	// Transmitter
	log_msg.Transmitter = raw_log_msg.Transmitter
	
	// Type
	log_msg.Type = raw_log_msg.Type
	
	// Data
	log_msg.Data = &raw_log_msg.Data
	
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
	j_data, err := json.Marshal(data)
	
	if err != nil {
		log.Println("A")
		return nil
	}
	
	data_str := string(j_data)
	raw_state_machine := NewRawStateMachine(data_str)
	
	if raw_state_machine == nil {
		log.Println("B")
		return nil
	}
	
	state_machine := NewStateMachine(raw_state_machine)
	
	if state_machine == nil {
		log.Println("C")
		return nil
	}
	
	return state_machine
}

func ParseStateTransitionMsg(data *map[string]interface{}) *TransitionEvent {
	j_data, err := json.Marshal(data)
	
	if err != nil {
		log.Println("A")
		return nil
	}
	
	data_str := string(j_data)
	raw_transition_event := NewRawTransitionEvent(data_str)
	
	if raw_transition_event == nil {
		log.Println("B")
		return nil
	}
	
	transition_event := NewTransitionEvent(raw_transition_event)
	
	if transition_event == nil {
		log.Println("C")
		return nil
	}
	
	return transition_event
}





