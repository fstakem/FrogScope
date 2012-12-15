package state_machine

import (

)


type State struct {
	Name string
	Id int
	EnterTimestamp []int64
	ExitTimestamp []int64
}

func(state State) Test() string {
	return ""
}

