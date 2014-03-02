package m

import (
	"encoding/json"
	"log"
)

type M map[string]interface{}

func (d M) JSON() []byte {
	s, err := json.Marshal(d)
	if err != nil {
		log.Println("[ERROR]: Failed to JSON encode err:" + err.Error())
	}
	return s
}

func (m M) ToM() M {
	return m
}
