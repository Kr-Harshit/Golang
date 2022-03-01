package handlers

import (
	"encoding/json"
	"io"
)

// Generic Error is generic Error Message returned by Server
type GenericError struct {
	Message string `json:"message"`
}

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(w io.Writer, i interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
