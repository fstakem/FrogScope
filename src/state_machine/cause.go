package state_machine

import (
	"bytes"
)

type Cause struct {
	Expression 	string
}

func NewCause(raw_cause *RawCause) *Cause {
	if raw_cause == nil {
		return nil
	}
	
	var cause Cause
	cause.Expression = raw_cause.Expression
	
	return &cause
}

func (this *Cause) Cmp(other *Cause) (equ bool) {
	if this.Expression != other.Expression {
		return false
	}

	return true
}

func (this *Cause) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Expression: " + this.Expression)
	
	return buffer.String()
}

