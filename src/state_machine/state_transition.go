package state_machine

import (

)

type StateTransistion struct {
	Name string
	To State
	From State
}

