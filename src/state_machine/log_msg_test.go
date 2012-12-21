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
	
	// Import file
	filename := "../../data/state_machine_1.json"
	log_msg := getRawMsg(t, filename)
	
	// Create state machine from imported file
	
	// Test to make sure the objects are the same
}

func TestParseStateTransitionMsg(t *testing.T) {
	// Test data
	test_transition_event := TransitionEvent{}
	test_transition_event.Id = 24
	test_transition_event.StateMachineId = 1
	test_transition_event.FromStateId = 4
	test_transition_event.ToStateId = 5
	test_transition_event.Timestamp = 123456789
	test_transition_event.Cause = make(map[string]string)
	test_transition_event.Cause["a"] = "10"
	
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







