// Code generated by go-swagger; DO NOT EDIT.

package category_service

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewAddRootCategoryParams creates a new AddRootCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRootCategoryParams() *AddRootCategoryParams {
	return &AddRootCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRootCategoryParamsWithTimeout creates a new AddRootCategoryParams object
// with the ability to set a timeout on a request.
func NewAddRootCategoryParamsWithTimeout(timeout time.Duration) *AddRootCategoryParams {
	return &AddRootCategoryParams{
		timeout: timeout,
	}
}

// NewAddRootCategoryParamsWithContext creates a new AddRootCategoryParams object
// with the ability to set a context for a request.
func NewAddRootCategoryParamsWithContext(ctx context.Context) *AddRootCategoryParams {
	return &AddRootCategoryParams{
		Context: ctx,
	}
}

// NewAddRootCategoryParamsWithHTTPClient creates a new AddRootCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRootCategoryParamsWithHTTPClient(client *http.Client) *AddRootCategoryParams {
	return &AddRootCategoryParams{
		HTTPClient: client,
	}
}

/*
AddRootCategoryParams contains all the parameters to send to the API endpoint

	for the add root category operation.

	Typically these are written to a http.Request.
*/
type AddRootCategoryParams struct {

	// Body.
	Body *models.CategoryContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add root category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRootCategoryParams) WithDefaults() *AddRootCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add root category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRootCategoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add root category params
func (o *AddRootCategoryParams) WithTimeout(timeout time.Duration) *AddRootCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add root category params
func (o *AddRootCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add root category params
func (o *AddRootCategoryParams) WithContext(ctx context.Context) *AddRootCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add root category params
func (o *AddRootCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add root category params
func (o *AddRootCategoryParams) WithHTTPClient(client *http.Client) *AddRootCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add root category params
func (o *AddRootCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add root category params
func (o *AddRootCategoryParams) WithBody(body *models.CategoryContext) *AddRootCategoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add root category params
func (o *AddRootCategoryParams) SetBody(body *models.CategoryContext) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRootCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
