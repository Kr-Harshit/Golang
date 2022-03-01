package handlers

import (
	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
)

// NOTE: Types defined here are purely for documentation Purposes
// these types are not used by handlers

// Generic Error message returned as string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation Error defined as an array of string
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Description of the error
	// in: body
	Body ValidationError
}

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products
	// in: body
	Body models.Products
}

// A Data Structure representing a single Product
// swagger:response productResponse
type productResponseWrapper struct {
	// One Product
	// in: body
	Body models.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

//swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to update or create
	// Note: the id field is ignored by create and update product
	// in: body
	// required: true
	Body models.Product
}

//swagger:parameters listSingleProdcut deleteProduct
type productIDParamsWrapper struct {
	// The id of Product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

//swagger:parameters listSingleProdcut listProducts
type productQueryParams struct {
	// Currency used when returning the price of product
	// when not specified currency is returned in USD
	// in: query
	// required: false
	Currency string `json:"currency"`
}
