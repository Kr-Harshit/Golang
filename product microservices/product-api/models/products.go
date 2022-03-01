package models

import (
	"context"
	"fmt"
	"time"

	protos "github.com/Kr-Harshit/golang-example/product-microservices/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

// ErrProductNotFound is a an error raised when a product can not be found in database
var ErrProductNotFound = fmt.Errorf("Product not found")

//  Products defines a slice of Product
type Products []*Product

// ProductsDB is database for products and currencyClient connection
type ProductsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
}

// NewProductsDB returns new ProductsDB
func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	return &ProductsDB{currency: c, log: l}
}

// getRate fetchs converted currency rate factor from currency gRPC server
// base currency is USD
// dest is the currency to which the rates are to be converted into
func (p *ProductsDB) getRate(dest string) (float64, error) {
	if dest == "" {
		return 1, nil
	}
	req := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["USD"]),
		Destination: protos.Currencies(protos.Currencies_value[dest]),
	}

	res, err := p.currency.GetRate(context.Background(), req)
	return res.Rate, err
}

// GetProducts returns all products from database
// currency is the currency to which the rates are to be converted into
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	rate, err := p.getRate(currency)

	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	modifiedProductList := Products{}
	for _, p := range productList {
		newProduct := *p
		newProduct.Price = newProduct.Price * rate
		modifiedProductList = append(modifiedProductList, &newProduct)
	}
	return modifiedProductList, nil
}

// GetProductByID returns a single product
// which matches the id from the database
// If a product with the given id does not exist in database
// the function returns ProductNotFound error
// currency is the currency to which the rates are to be converted into
func (p *ProductsDB) GetProductByID(id int, currency string) (*Product, error) {
	i := p.findIndexByProductID(id)
	if i == -1 {
		return nil, ErrProductNotFound
	}

	rate, err := p.getRate(currency)
	p.log.Info("rate", rate)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	newProduct := *productList[i]
	newProduct.Price = newProduct.Price * rate
	return &newProduct, nil
}

// AddProduct adds a new product to database
func (p *ProductsDB) AddProduct(pr *Product) {
	pr.ID = productList[len(productList)-1].ID + 1
	pr.CreatedOn = time.Now().UTC().String()
	pr.UpdatedOn = time.Now().UTC().String()
	productList = append(productList, pr)
}

// UpdateProduct replaces a product in  the database with the given item.
// If a product with the given id does not exist in the database
// this function returns a  ProductNotFound error
func (p *ProductsDB) UpdateProduct(pr *Product) error {
	i := p.findIndexByProductID(pr.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	productList[i] = pr
	return nil
}

// DeleteProduct deletes a product from database
func (p *ProductsDB) DeleteProduct(id int) error {
	i := p.findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

//findIndexByProductID finds the index of a product in the database
// return -1 when no product can be found
func (p *ProductsDB) findIndexByProductID(id int) int {
	for i, pr := range productList {
		if pr.ID == id {
			return i
		}
	}
	return -1
}

// productList is a silce of Product
// Temporary database
var productList = Products{
	&Product{
		ID:          1,
		Name:        "latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "short strong coffee without milk",
		Price:       1.99,
		SKU:         "fdj34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "frappe",
		Description: "Don't know what it is?",
		Price:       1.25,
		SKU:         "kd345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
