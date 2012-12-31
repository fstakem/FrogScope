package sequence_diagram

import (

)

type SequenceEvent struct {
	Name			string
	Id			int
	SeqDiagId	int
	FromId		int
	ToId			int
	Timestamp	uint64
	LinkId		int
}

func NewSequenceEvent(raw_sequence_event *RawSequenceEvent) (sequence_event *SequenceEvent) {
	if raw_sequence_event == nil {
		return nil
	}
	
	var sequence_event SequenceEvent
	
	
	
	return &sequence_event
}

func (this *SequenceEvent) Cmp(other *SequenceEvent) (equ bool) {

}

func (this *SequenceEvent) String() (output string) {

}

