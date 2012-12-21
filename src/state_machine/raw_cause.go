package state_machine

import (
	"encoding/json"
	"bytes"
)

type RawCause struct {
	Symbol 	string
	Value	string
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
	if this.Symbol != other.Symbol {
		return false
	} else if this.Value != other.Value {
		return false
	}
	
	return true
}

func (this *RawCause) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Symbol: " + this.Symbol)
	buffer.WriteString("  Value: " + this.Value)
	
	return buffer.String()
}

