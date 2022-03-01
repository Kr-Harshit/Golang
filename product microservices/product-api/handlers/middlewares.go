package handlers

import (
	"context"
	"net/http"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/models"
)

// ValidateProduct is a middleWare that validates the product in the request and calls next if ok
func (p *Products) ValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &models.Product{}

		if err := models.FromJSON(r.Body, prod); err != nil {
			p.l.Error("Deserializing product from request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			models.ToJSON(w, &GenericError{Message: err.Error()})
			return
		}

		// Validate the Product
		if errs := p.v.Validate(prod); len(errs) != 0 {
			p.l.Error("Unable to validate product request", "error", errs)
			w.WriteHeader(http.StatusBadRequest)
			models.ToJSON(w, &ValidationError{Messages: errs.Errors()})
			return
		}

		// add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call next handler, which can be next middleware in chain or final Handler.
		next.ServeHTTP(w, r)
	})
}
