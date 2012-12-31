package sequence_diagram

import (
	"strconv"
	"log"
)

type Component struct {
	Name		string
	Id		int
}

func NewComponent(raw_component *RawComponent) (component *Component) {
	if raw_component == nil {
		return nil
	}
	
	var component Component
	
	// Name
	component.Name = raw_component.Name
	
	// Id
	id, err := strconv.Atoi(raw_component.Id)
	if err != nil {
		log.Println("RawComponent.Id not properly parsed.")
		return nil
	}
	component.Id = id
	
	return &component
}

func (this *Component) Cmp(other *Component) (equ bool) {
	if this.Name != other.Name {
		return false
	} else if this.Id != other.Id {
		return false
	}
	
	return true
}

func (this *Component) String() (output string) {
	var buffer bytes.Buffer
	buffer.WriteString("Name: " + this.Name)
	buffer.WriteString(" Id: " + strconv.Itoa(this.Id))
	
	return buffer.String()
}

