package state_machine

import (
	"io/ioutil"
)

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
    if err != nil { return nil, err }
    
    return data, nil
}

