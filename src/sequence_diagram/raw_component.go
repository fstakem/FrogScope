package sequence_diagram

import (
	"encoding/json"
	"bytes"
)

type RawComponent struct {
	Name 	string
	Id		string
}

func NewRawComponent(msg *string) (raw_component *RawComponent) {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_component RawComponent
	err := json.Unmarshal([]byte msg, &raw_component)
	
	if err != nil {
		return nil
	}
	
	return &raw_component
}

func (this *RawComponent) Cmp(other *RawComponent) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	}
	
	return true
}

func (this *RawComponent) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + this.Id)
	
	return buffer.String()
} 
