package state_machine

import (
	"encoding/json"
	"bytes"
)

type RawLogMsg struct {
	Protocol 	string 				
	Version 		string					
	SeqNum 		string				
	Scope 		string				
	Timestamp 	string				
	Transmitter string				
	Type 		string				
	Data 		map[string]interface{}	
}

func NewRawLogMsg(msg string) *RawLogMsg {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_log_msg RawLogMsg
	err := json.Unmarshal([]byte(msg), &raw_log_msg)
	
	if err != nil {
		return nil
	}
	
	return &raw_log_msg
}

func (this *RawLogMsg) Cmp(other *RawLogMsg) (equ bool) {
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
	}  else if this.Transmitter != other.Transmitter {
		return false
	} else if this.Type != other.Type {
		return false
	}
	
	return true
}

func (this *RawLogMsg) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Protocol: " + this.Protocol)
	buffer.WriteString("  Version: " + this.Version)
	buffer.WriteString("  SeqNum: " + this.SeqNum)
	buffer.WriteString("  Scope: " + this.Scope)
	buffer.WriteString("  Timestamp: " + this.Timestamp)
	buffer.WriteString("  Transmitter: " + this.Transmitter)
	buffer.WriteString("  Type: " + this.Type)
	
	return buffer.String()
}


