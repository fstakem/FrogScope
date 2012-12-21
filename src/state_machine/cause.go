package state_machine

import (

)

type Cause struct {
	Symbol 	string
	Value	string
}

func NewCause(raw_cause *RawCause) *Cause {
	if raw_cause == nil {
		return nil
	}
	
	var cause Cause
	
	return &cause
}

