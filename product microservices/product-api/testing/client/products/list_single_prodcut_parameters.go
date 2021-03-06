// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListSingleProdcutParams creates a new ListSingleProdcutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSingleProdcutParams() *ListSingleProdcutParams {
	return &ListSingleProdcutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSingleProdcutParamsWithTimeout creates a new ListSingleProdcutParams object
// with the ability to set a timeout on a request.
func NewListSingleProdcutParamsWithTimeout(timeout time.Duration) *ListSingleProdcutParams {
	return &ListSingleProdcutParams{
		timeout: timeout,
	}
}

// NewListSingleProdcutParamsWithContext creates a new ListSingleProdcutParams object
// with the ability to set a context for a request.
func NewListSingleProdcutParamsWithContext(ctx context.Context) *ListSingleProdcutParams {
	return &ListSingleProdcutParams{
		Context: ctx,
	}
}

// NewListSingleProdcutParamsWithHTTPClient creates a new ListSingleProdcutParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSingleProdcutParamsWithHTTPClient(client *http.Client) *ListSingleProdcutParams {
	return &ListSingleProdcutParams{
		HTTPClient: client,
	}
}

/* ListSingleProdcutParams contains all the parameters to send to the API endpoint
   for the list single prodcut operation.

   Typically these are written to a http.Request.
*/
type ListSingleProdcutParams struct {

	/* ID.

	   The id of Product for which the operation relates

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list single prodcut params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSingleProdcutParams) WithDefaults() *ListSingleProdcutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list single prodcut params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSingleProdcutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list single prodcut params
func (o *ListSingleProdcutParams) WithTimeout(timeout time.Duration) *ListSingleProdcutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list single prodcut params
func (o *ListSingleProdcutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list single prodcut params
func (o *ListSingleProdcutParams) WithContext(ctx context.Context) *ListSingleProdcutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list single prodcut params
func (o *ListSingleProdcutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list single prodcut params
func (o *ListSingleProdcutParams) WithHTTPClient(client *http.Client) *ListSingleProdcutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list single prodcut params
func (o *ListSingleProdcutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list single prodcut params
func (o *ListSingleProdcutParams) WithID(id int64) *ListSingleProdcutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list single prodcut params
func (o *ListSingleProdcutParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListSingleProdcutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
