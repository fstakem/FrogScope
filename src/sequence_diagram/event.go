package sequence_diagram

import (

)

type Event struct {
	Name		string
	Id		int
	FromId	int
	ToId		int
	LinkId	int
}

func NewEvent(raw_event *RawEvent) (event *Event) {
	if raw_event == nil {
		return nil
	}
	
	var event Event
	
	// Name
	event.Name = raw_event.Name
	
	// Id
	id, err := strconv.Atoi(raw_event.Id)
	if err != nil {
		log.Println("RawEvent.Id not properly parsed.")
		return nil
	}
	event.Id = id
	
	// FromId
	from_id, err := strconv.Atoi(raw_event.FromId)
	if err != nil {
		log.Println("RawEvent.FromId not properly parsed.")
		return nil
	}
	event.FromId = from_id
	
	// ToId
	to_id, err := strconv.Atoi(raw_event.ToId)
	if err != nil {
		log.Println("RawEvent.ToId not properly parsed.")
		return nil
	}
	event.ToId = to_id
	
	// LinkId
	link_id, err := strconv.Atoi(raw_event.LinkId)
	if err != nil {
		log.Println("RawEvent.LinkId not properly parsed.")
		return nil
	}
	event.LinkId = link_id
	
	return &event
}

func (this *Event) Cmp(other *Event) (equ bool) {
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

func (this *Event) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString("  Id: " + strconv.Itoa(this.Id))
	buffer.WriteString("  FromId: " + strconv.Itoa(this.FromId))
	buffer.WriteString("  ToId: " + strconv.Itoa(this.ToId))
	buffer.WriteString("  LinkId: " + strconv.Itoa(this.LinkId))
	
	return buffer.String()
}


