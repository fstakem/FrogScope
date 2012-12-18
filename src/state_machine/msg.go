package state_machine

import (
	"fmt"
	"encoding/json"
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
	Protocol string
	Version int
	SeqNum int
	Scope string
	Timestamp int64
	Transmitter string
	Type string
	Data map[string]string
}

func NewLogMsg(msg string) *LogMsg {
	if len(msg) == 0 {
		return nil
	}
	
	var log_msg LogMsg
	err := json.Unmarshal([]byte(msg), &log_msg)
	
	if err == nil {
		return nil
	}
	
	return &log_msg
}

func (this *LogMsg) Cmp(other *LogMsg) (equ bool) {
	if this.Protocol != other.Protocol { 
		fmt.Println(this.String())
		fmt.Println(other.String())
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
	buffer.WriteString("  Timestamp: " + strconv.FormatInt(this.Timestamp, 10))
	buffer.WriteString("  Transmitter: " + this.Transmitter)
	buffer.WriteString("  Type: " + this.Type)
	
	return buffer.String()
}
 
func ParseStateMachineMsg(data *map[string]string) *StateMachine {

	for key, value := range *data {
		fmt.Println("%s::%s",  key, value)
	}
	
	return nil
}

func ParseStateTransitionMsg(data *map[string]string) *StateTransition {
	return nil
}




