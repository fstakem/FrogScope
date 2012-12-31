package sequence_diagram

import (
	"encoding/json"
	"bytes"
)

type RawSequenceDiagram struct {
	Name			string
	Id			string
	Components	[]map[string]interface{}
	Events		[]map[string]interface{}
}

func NewRawSequenceDiagram(msg string) (raw_sequence_diagram *RawSequenceDiagram) {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_sequence_diagram RawSequenceDiagram
	
	err := json.Unmarshal([]byte msg, &raw_sequence_diagram)
	
	if err != nil {
		return nil
	}
	
	return &raw_sequence_diagram
}

func (this *RawSequenceDiagram) Cmp(other *RawSequenceDiagram) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	}
	
	return true
}

func (this *RawSequenceDiagram) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + this.Id)
	
	return buffer.String()
}

