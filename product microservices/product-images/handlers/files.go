// Package  classification of Product-Images API
//
// Documentation for Product-Images API
//
//  Schemes: http
//  BasePath: /
//  Port: 9090
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//  swagger:meta

package handlers

import (
	"log"
	"net/http"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-images/files"
	"github.com/gorilla/mux"
)

//Files is a handler for reading and writing files
type Files struct {
	logger *log.Logger
	store  files.Storage
}

// NewFiles create a new File handler
func NewFiles(s files.Storage, l *log.Logger) *Files {
	return &Files{logger: l, store: s}
}

// getPATH extracts id and filename from request URI
func (f *Files) getPATH(r *http.Request) (string, string) {
	vars := mux.Vars(r)
	id := vars["id"]
	filename := vars["filename"]

	return id, filename
}
