package state_machine

import (
	//"json"
)

type StateMachine struct {
	Name string
	Id int
	States []State
	Transitions []StateTransition
}

func(this *StateMachine) AnalyzeLog(log string) {

}

