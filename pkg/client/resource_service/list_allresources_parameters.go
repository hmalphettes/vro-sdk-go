// Code generated by go-swagger; DO NOT EDIT.

package resource_service

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
)

// NewListAllresourcesParams creates a new ListAllresourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAllresourcesParams() *ListAllresourcesParams {
	return &ListAllresourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAllresourcesParamsWithTimeout creates a new ListAllresourcesParams object
// with the ability to set a timeout on a request.
func NewListAllresourcesParamsWithTimeout(timeout time.Duration) *ListAllresourcesParams {
	return &ListAllresourcesParams{
		timeout: timeout,
	}
}

// NewListAllresourcesParamsWithContext creates a new ListAllresourcesParams object
// with the ability to set a context for a request.
func NewListAllresourcesParamsWithContext(ctx context.Context) *ListAllresourcesParams {
	return &ListAllresourcesParams{
		Context: ctx,
	}
}

// NewListAllresourcesParamsWithHTTPClient creates a new ListAllresourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAllresourcesParamsWithHTTPClient(client *http.Client) *ListAllresourcesParams {
	return &ListAllresourcesParams{
		HTTPClient: client,
	}
}

/*
ListAllresourcesParams contains all the parameters to send to the API endpoint

	for the list allresources operation.

	Typically these are written to a http.Request.
*/
type ListAllresourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list allresources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAllresourcesParams) WithDefaults() *ListAllresourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list allresources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAllresourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list allresources params
func (o *ListAllresourcesParams) WithTimeout(timeout time.Duration) *ListAllresourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list allresources params
func (o *ListAllresourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list allresources params
func (o *ListAllresourcesParams) WithContext(ctx context.Context) *ListAllresourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list allresources params
func (o *ListAllresourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list allresources params
func (o *ListAllresourcesParams) WithHTTPClient(client *http.Client) *ListAllresourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list allresources params
func (o *ListAllresourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllresourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
