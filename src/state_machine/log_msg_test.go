package state_machine

import (
	"io/ioutil"
	"testing"
	//"encoding/json"
)

func TestNewLogMsg(t *testing.T) {
	// Test data
	test_log_msg := LogMsg{}
	test_log_msg.Protocol = "frogscope"
	test_log_msg.Version = 1
	test_log_msg.SeqNum = 1
	test_log_msg.Scope = "test"
	test_log_msg.Timestamp = 123456789
	test_log_msg.Transmitter = "main"
	test_log_msg.Type = "state_transition"
	
	// Import file
	filename := "../../data/state_transition_event_1.json"
	log_msg := getRawMsg(t, filename)
	
	// Test to make sure objects are the same
	if !log_msg.Cmp(&test_log_msg) {
		t.Log("The JSON data was not interpreted correctly.")
		t.Fail()
	}
}

func TestParseStateMachineMsg(t *testing.T) { 
	// Test data
	test_state_machine := StateMachine{}
	test_state_machine.Name = "machina"
	test_state_machine.Id = 24
	
	init_state := State{}
	init_state.Name = "init"
	init_state.Id = 1
	test_state_machine.States = append(test_state_machine.States, init_state)
	
	sec_state := State{}
	sec_state.Name = "second"
	sec_state.Id = 2
	test_state_machine.States = append(test_state_machine.States, sec_state)
	
	state_transition := StateTransition{}
	state_transition.Name = "x"
	state_transition.Id = 1
	state_transition.From = init_state
	state_transition.To = sec_state
	cause := Cause{}
	cause.Expression = "a = 3"
	state_transition.Causes = append(state_transition.Causes, cause)
	cause = Cause{}
	cause.Expression = "b > 4"
	state_transition.Causes = append(state_transition.Causes, cause)
	test_state_machine.Transitions = append(test_state_machine.Transitions, state_transition)
	
	state_transition = StateTransition{}
	state_transition.Name = "y"
	state_transition.Id = 2
	state_transition.From = sec_state
	state_transition.To = init_state
	cause = Cause{}
	cause.Expression = "a = 5"
	state_transition.Causes = append(state_transition.Causes, cause)
	cause = Cause{}
	cause.Expression = "b > 6"
	state_transition.Causes = append(state_transition.Causes, cause)
	test_state_machine.Transitions = append(test_state_machine.Transitions, state_transition)
	
	// Import file
	filename := "../../data/state_machine_1.json"
	log_msg := getRawMsg(t, filename)
	
	// Create state machine from imported file
	state_machine := ParseStateMachineMsg(log_msg.Data)
	t.Log("State machine imported from the file:")
	t.Logf(state_machine.String())
	
	// Test to make sure the objects are the same
	if !state_machine.Cmp(&test_state_machine) {
		t.Log("The JSON data was not interpreted correctly.")
		t.Fail()
	}
	
}

func TestParseStateTransitionMsg(t *testing.T) {
	// Test data
	test_transition_event := TransitionEvent{}
	test_transition_event.Id = 24
	test_transition_event.StateMachineId = 1
	test_transition_event.FromStateId = 4
	test_transition_event.ToStateId = 5
	test_transition_event.Timestamp = 123456789
	
	cause := Cause{}
	cause.Expression = "a = 10"
	test_transition_event.Causes = append(test_transition_event.Causes, cause)
	
	cause = Cause{}
	cause.Expression = "b < 12"
	test_transition_event.Causes = append(test_transition_event.Causes, cause)
	
	// Import file
	filename := "../../data/state_transition_event_1.json"
	log_msg := getRawMsg(t, filename)
	
	// Create transition event from imported file
	transition_event := ParseStateTransitionMsg(log_msg.Data)
	t.Log("Transition event imported from the file:")
	t.Logf(transition_event.String())
	
	// Test to make sure objects are the same
	if !transition_event.Cmp(&test_transition_event) {
		t.Log("The JSON data was not interpreted correctly.")
		t.Fail()
	}
}

func readFile(filename string) (data string, err error) {
	raw_data, err := ioutil.ReadFile(filename)
    if err != nil { return "", err }
    
    return string(raw_data), nil
}

func getRawMsg(t *testing.T, filename string) (log_msg *LogMsg) {
	// Create msg from imported file
	raw_data, err := readFile(filename)
	if err != nil {
		t.Log("Unable to open file.")
		t.Fail()
	}
	t.Log("Message imported from the file:")
	t.Log(raw_data)
	
	// Create data structure from data
	raw_log_msg := NewRawLogMsg(raw_data)
	if raw_log_msg == nil {
		t.Log("Unable to parse JSON data.")
		t.Fail()
	}
	
	// Transform data structure to new one
	log_msg = NewLogMsg(raw_log_msg)
	if log_msg == nil {
		t.Log("Unable to parse raw data.")
		t.Fail()
	}
	t.Log("Message created from the data:")
	t.Log(log_msg.String())
	
	return log_msg
}







