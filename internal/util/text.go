package util

import (
	"encoding/json"
	"log"
)

type Jsonable interface {
	ToJson() string
}

func MarshalOrPanic(a any) string {
	b, err := json.Marshal(a)
	if err != nil {
		log.Panicf("error marshalling: %v", err)
	}
	return string(b)
}
