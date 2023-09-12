package util

import (
	"encoding/json"
	"log"
)

func MarshalOrPanic(a any) string {
	b, err := json.Marshal(a)
	if err != nil {
		log.Panicf("error marshalling: %v", err)
	}
	return string(b)
}
