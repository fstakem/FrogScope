package state_machine

import (
	"bytes"
	"strconv"
	"log"
)


type State struct {
	Name 			string
	Id 				int
	EnterTimestamps 	[]uint64
	ExitTimestamps 	[]uint64
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
	
	if len(this.EnterTimestamps) == len(other.EnterTimestamps) {
		for i := range this.EnterTimestamps {
			if this.EnterTimestamps[i] != other.EnterTimestamps[i] {
				return false
			}
		}
	} else {
		return false
	}
	
	if len(this.ExitTimestamps) == len(other.ExitTimestamps) {
		for i := range this.ExitTimestamps {
			if this.ExitTimestamps[i] != other.ExitTimestamps[i] {
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
	for i := range this.EnterTimestamps {
		buffer.WriteString(strconv.FormatUint(this.EnterTimestamps[i], 10))
		buffer.WriteString(" ")
	}
	
	buffer.WriteString(" Exit Timestamps: ")
	for i := range this.ExitTimestamps {
		buffer.WriteString(strconv.FormatUint(this.ExitTimestamps[i], 10))
		buffer.WriteString(" ")
	}
	
	return buffer.String()
}
