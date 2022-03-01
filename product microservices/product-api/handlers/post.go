package handlers

import (
	"net/http"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//  	201: noContentResponse
// 400: errorResponse
// 422: errorValidation
// 501: errorResponse

// Create handles POST request to add New products
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(*models.Product)

	p.l.Debug("Inserting product", "product", prod)
	p.productDB.AddProduct(prod)
	w.WriteHeader(http.StatusNoContent)
}
