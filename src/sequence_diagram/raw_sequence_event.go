package sequence_diagram

import (
	"encoding/json"
	"bytes"
)

type RawSequenceEvent struct {
	Name			string
	Id			string
	SeqDiagId	string
	FromId		string
	ToId			string
	Timestamp	string
	LinkId		string
}

func NewRawSequenceEvent(msg string) (raw_sequence_event *RawSequenceEvent) {
	if len(msg) == 0 {
		return nil
	}
	
	var raw_sequence_event RawSequenceEvent
	
	err := json.Unmarshal([]byte msg, &raw_sequence_event)
	
	if err != nil {
		return nil
	}
	
	return &raw_sequence_event
}

func (this *RawSequenceEvent) Cmp(other *RawSequenceEvent) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	} else if this.SeqDiagId != other.SeqDiagId {
		return false
	} else if this.FromId != other.FromId {
		return false
	} else if this.ToId	 != other.ToId {
		return false
	} else if this.Timestamp != other.Timestamp {
		return false
	} else if this.LinkId != other.LinkId {
		return false
	}
	
	return true
}

func (this *RawSequenceEvent) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + this.Id)
	buffer.WriteString("  SeqDiagId: " + this.SeqDiagId)
	buffer.WriteString("  FromId: " + this.FromId)
	buffer.WriteString("  ToId: " + this.ToId)
	buffer.WriteString("  Timestamp: " + this.Timestamp)
	buffer.WriteString("  LinkId: " + this.LinkId)
	
	return buffer.String()
}

