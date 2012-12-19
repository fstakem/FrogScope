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

}

func TestParseStateTransitionMsg(t *testing.T) {
	// Parse the state transition specific data
	//state_transition := ParseStateTransitionMsg(&log_msg.Data)
	//data, err := json.Marshal(&log_msg.Data)
	//data_str := string(data)
	//t.Log(data_str)
	//transition_event := NewTransitionEvent(data_str)
	//t.Log(transition_event.String())
}

func readFile(filename string) (data string, err error) {
	raw_data, err := ioutil.ReadFile(filename)
    if err != nil { return "", err }
    
    return string(raw_data), nil
}

func getRawMsg(t *testing.T, filename string) (log_msg *LogMsg) {
	raw_data, err := readFile(filename)
	if err != nil {
		t.Log("Unable to open file.")
		t.Fail()
	}
	
	// Create object from imported file
	t.Log("Data imported from the file:")
	t.Log(raw_data)
	raw_log_msg := NewRawLogMsg(raw_data)
	
	// Make sure an object was returned
	if raw_log_msg == nil {
		t.Log("Unable to parse JSON data.")
		t.Fail()
	}
	
	// Create internal object from json based object
	log_msg = NewLogMsg(raw_log_msg)
	
	// Make sure an object was returned
	if log_msg == nil {
		t.Log("Unable to parse raw data.")
		t.Fail()
	}
	
	return log_msg
}







