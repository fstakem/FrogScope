package main 

// Libraries
import (
	"fmt"
	"time"
	"./state_machine"
)

// Globals
// TODO

// Initialization
func init() {
	fmt.Println("Entering init function")
	
	var sm state_machine.Timestamp 
	sm.EnterTime = time.Now()
	time.Sleep(5)
	sm.ExitTime = time.Now()
	
	fmt.Println("Exiting init function")
}

// Main
func main() {
	fmt.Println("Entering main function")
	
	fmt.Println("Exiting main function")
}

// Other functions
func createStateMachine() state_machine.StateMachine {
	var sm state_machine.StateMachine
	
	// TODO
	
	return sm
}

