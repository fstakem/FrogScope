package state_machine

import (

)

type StateTransistion struct {
	Name string
	Id int
	timestamp int64
	From State
	To State
	Cause string
}

