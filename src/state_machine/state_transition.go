package state_machine

import (

)

type StateTransition struct {
	Name string
	Id int
	timestamp int64
	From State
	To State
	Cause string
}

