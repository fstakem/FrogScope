package sequence_diagram

import (
	"encoding/json"
	"bytes"
)

type RawEvent struct {
	Name		string
	Id		string
	FromId	string
	ToId		string
	LinkId	string
}

func NewRawEvent(string msg) (raw_event *RawEvent) {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_event RawEvent
	
	err := json.Unmarshal([]byte msg, &raw_event)
	
	if err != nil {
		return nil
	}
	
	return &raw_event
}

func (this *RawEvent) Cmp(other *RawEvent) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	} else if this.FromId != other.FromId {
		return false
	} else if this.ToId != other.ToId {
		return false
	} else if this.LinkId != other.LinkId {
		return false
	}
	
	return true
}

func (this *RawEvent) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + this.Id)
	buffer.WriteString("  FromId: " + this.FromId)
	buffer.WriteString("  ToId: " + this.ToId)
	buffer.WriteString("  LinkId: " + this.LinkId)
	
	return buffer.String()
}

