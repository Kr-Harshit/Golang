package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/testing/client"
	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/testing/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost(fmt.Sprintf("localhost%s", os.Getenv("TEST_CLIENT_PORT")))
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", prods)
	t.Fail()
}
