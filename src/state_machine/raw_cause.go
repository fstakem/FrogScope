package state_machine

import (
	"encoding/json"
	"bytes"
)

type RawCause struct {
	Expression 	string
}

func NewRawCause(msg string) *RawCause {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_cause RawCause
	err := json.Unmarshal([]byte(msg), &raw_cause)
	
	if err != nil {
		return nil
	}
	
	return &raw_cause
}

func (this *RawCause) Cmp(other *RawCause) (equ bool) {
	if this.Expression != other.Expression {
		return false
	} 
	
	return true
}

func (this *RawCause) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Expression: " + this.Expression)
	
	return buffer.String()
}

