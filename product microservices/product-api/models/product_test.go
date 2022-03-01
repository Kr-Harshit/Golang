package models

import "testing"

func TestValidateProduct(t *testing.T) {
	prod := &Product{Name: "fapacciuno", Price: 2.34, SKU: "abz-werd-wdsfg"}

	if valErr := prod.Validate(); valErr != nil {
		t.Fatal(valErr)
	}
}
