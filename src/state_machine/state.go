package state_machine

import (
	"bytes"
	"strconv"
	"log"
)


type State struct {
	Name 			string
	Id 				int
	EnterTimestamp 	[]uint64
	ExitTimestamp 	[]uint64
}

func NewState(raw_state *RawState) *State {
	if raw_state == nil {
		return nil
	}
	
	var state State
	
	// Name 
	state.Name = raw_state.Name
	
	// Id 
	id, err := strconv.Atoi(raw_state.Id)
	if err != nil {
		log.Println("RawState.Id not properly parsed.")
		return nil
	}
	state.Id = id
	
	return &state
}

func (this *State) Cmp(other *State) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	}
	
	if len(this.EnterTimestamp) == len(other.EnterTimestamp) {
		for i := range this.EnterTimestamp {
			if this.EnterTimestamp[i] != other.EnterTimestamp[i] {
				return false
			}
		}
	} else {
		return false
	}
	
	if len(this.ExitTimestamp) == len(other.ExitTimestamp) {
		for i := range this.ExitTimestamp {
			if this.ExitTimestamp[i] != other.ExitTimestamp[i] {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

func (this *State) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString(" Id: " + strconv.Itoa(this.Id))
	
	buffer.WriteString(" Enter Timestamps: ")
	for i := range this.EnterTimestamp {
		buffer.WriteString(strconv.FormatUint(this.EnterTimestamp[i], 10))
	}
	
	buffer.WriteString(" Exit Timestamps: ")
	for i := range this.ExitTimestamp {
		buffer.WriteString(strconv.FormatUint(this.ExitTimestamp[i], 10))
	}
	
	return buffer.String()
}
