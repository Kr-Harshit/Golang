// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Kr-Harshit/golang-example/product-microservices/product-api/testing/models"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProductBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProductInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProductCreated creates a UpdateProductCreated with default headers values
func NewUpdateProductCreated() *UpdateProductCreated {
	return &UpdateProductCreated{}
}

/* UpdateProductCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type UpdateProductCreated struct {
}

func (o *UpdateProductCreated) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductCreated ", 201)
}

func (o *UpdateProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductBadRequest creates a UpdateProductBadRequest with default headers values
func NewUpdateProductBadRequest() *UpdateProductBadRequest {
	return &UpdateProductBadRequest{}
}

/* UpdateProductBadRequest describes a response with status code 400, with default header values.

Validation Error defined as an array of string
*/
type UpdateProductBadRequest struct {
	Payload *models.ValidationError
}

func (o *UpdateProductBadRequest) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateProductBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateProductBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductNotFound creates a UpdateProductNotFound with default headers values
func NewUpdateProductNotFound() *UpdateProductNotFound {
	return &UpdateProductNotFound{}
}

/* UpdateProductNotFound describes a response with status code 404, with default header values.

Generic Error message returned as string
*/
type UpdateProductNotFound struct {
	Payload *models.GenericError
}

func (o *UpdateProductNotFound) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductNotFound  %+v", 404, o.Payload)
}
func (o *UpdateProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductInternalServerError creates a UpdateProductInternalServerError with default headers values
func NewUpdateProductInternalServerError() *UpdateProductInternalServerError {
	return &UpdateProductInternalServerError{}
}

/* UpdateProductInternalServerError describes a response with status code 500, with default header values.

Generic Error message returned as string
*/
type UpdateProductInternalServerError struct {
	Payload *models.GenericError
}

func (o *UpdateProductInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateProductInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateProductInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}