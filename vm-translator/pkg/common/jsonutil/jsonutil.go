package jsonutil

import (
	"encoding/json"
	"log"
)

// Marshal .
func Marshal(obj interface{}) string {
	v, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(v)
}

// MarshalPrint .
func Print(obj interface{}) {
	str := Marshal(obj)
	log.Println("obj is ", str)
}
