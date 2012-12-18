package state_machine

import (
	"io/ioutil"
	"testing"
)

func readFile(filename string) (data string, err error) {
	var raw_data []byte
	raw_data, err = ioutil.ReadFile(filename)
    if err != nil { return "", err }
    
    return string(raw_data), nil
}

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
	test_log_msg.Data = make(map[string]string)
	
	// Import file
	filename := "../../data/state_transition_event_1.json"
	raw_data, err := readFile(filename)
	if err != nil {
		t.Log("Unable to open file.")
		t.Fail()
	}
	
	// Create object from imported file
	t.Log("Data imported from the file:")
	t.Log(raw_data)
	log_msg := NewLogMsg(raw_data)
	
	// Test to make sure objects are the same
	if !log_msg.Cmp(&test_log_msg) {
		t.Log("The JSON data was not interpretted correctly.")
		t.Fail()
	}
}

func TestParseStateMachineMsg(t *testing.T) { 

}

func TestParseStateTransitionMsg(t *testing.T) {
	
}

