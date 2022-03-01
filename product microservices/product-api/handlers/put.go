package handlers

import (
	"net/http"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
)

// swagger:route PUT /products products updateProduct
// Update a product details
// responses:
//  201: noContentResponse
// 	400: errorValidation
// 	404: errorResponse
//  500: errorResponse

// Update handles PUT request to update products
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {

	// fetch the product from context
	prod := r.Context().Value(KeyProduct{}).(*models.Product)
	p.l.Debug("Update record", "id", prod.ID)

	err := p.productDB.UpdateProduct(prod)
	switch err {
	case nil:
	case models.ErrProductNotFound:
		p.l.Error("Product not found", "id", prod.ID, "error", err)
		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(w, &GenericError{Message: err.Error()})
		return
	default:
		p.l.Error("Unable to update product", "id", prod.ID, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(w, &GenericError{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
