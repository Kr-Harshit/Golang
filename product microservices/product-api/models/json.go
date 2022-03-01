package models

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(w io.Writer, i interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the objects from JSON string
// in an io.Reader to given interface
func FromJSON(r io.Reader, i interface{}) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}
