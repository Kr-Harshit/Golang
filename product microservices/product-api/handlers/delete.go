package handlers

import (
	"net/http"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product in database
//
// responses:
//  	201: noContentResponse
// 404: errorResponse
// 501: errorResponse

// DeleteProduct handles deletes request and remove an item from database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Debug("Delete record id", id)

	err := p.productDB.DeleteProduct(id)
	switch err {
	case nil:
	case models.ErrProductNotFound:
		p.l.Error("Record does not exist", "id", id)
		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(w, &GenericError{Message: err.Error()})
		return
	default:
		p.l.Error("Unable to delete record", "id", id, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(w, &GenericError{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
