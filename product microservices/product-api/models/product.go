package models

//Product defines the structure for an API product
//swagger:model
type Product struct {
	// ID of product generated by system
	//
	// required: false
	// min: 1
	// readOnly: true
	ID int `json:"_id"` // unique identifer for product

	// name of the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// description of product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// price of product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"required,gt=0"`

	// SKU of product
	//
	// required: true
	// Pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`

	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}
