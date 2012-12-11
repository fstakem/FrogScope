package state_machine

import (

)

type StateTransistion struct {
	Name string
	Id int
	Timestamps []Timestamp
	To State
	From State
}

