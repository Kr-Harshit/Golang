// Package  classification of Product API
//
// Documentation for Product API
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
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
	"github.com/hashicorp/go-hclog"

	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Products Handler for getting and updating products
type Products struct {
	l         hclog.Logger
	v         *models.Validation
	productDB *models.ProductsDB
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l hclog.Logger, v *models.Validation, p *models.ProductsDB) *Products {
	return &Products{l, v, p}
}

// ErrInvalidProductPath is an error message  when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of vallidation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the Product ID from the URI
// Panics if it cannot convert the id  into an integer

func getProductID(r *http.Request) int {
	// parse the id from the url
	vars := mux.Vars(r)

	//convert id to integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//  this should never happen
		// as router ensures that id is avalid number
		panic(err)
	}
	return id
}
