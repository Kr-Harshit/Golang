package handlers

import (
	"net/http"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
)

// swagger:route GET /products products listProducts
// Returns a list of Products from the database
//
// responses:
//  200: productsResponse

// GetProducts handles GET requests and returns all current products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Debug("Get all records")
	w.Header().Add("Content-Type", "application/json")

	cur := r.URL.Query().Get("currency")

	prods, err := p.productDB.GetProducts(cur)
	if err != nil {
		p.l.Error("fetching Product from database", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(w, &GenericError{Message: err.Error()})
		return
	}

	if err := models.ToJSON(w, prods); err != nil {
		p.l.Error("Serializing Product", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(w, &GenericError{Message: err.Error()})
	}
}

// swagger:route GET /products/{id} products listSingleProdcut
// Returns a specific product from database
//
// responses:
//  200: productResponse
//	404: errorResponse

// GetProduct handles GET requests and returns a single product of given id
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Debug("Get record", "id", id)
	cur := r.URL.Query().Get("currency")

	prod, err := p.productDB.GetProductByID(id, cur)

	switch err {
	case nil:
	case models.ErrProductNotFound:
		p.l.Error("Product not found", "error", err)
		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(w, &GenericError{Message: err.Error()})
		return
	default:
		p.l.Error("Unable to fetch Product", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(w, &GenericError{Message: err.Error()})
		return
	}

	if err := models.ToJSON(w, prod); err != nil {
		p.l.Error("Serializing Product", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(w, &GenericError{Message: err.Error()})
	}
}
