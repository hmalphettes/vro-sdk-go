// Code generated by go-swagger; DO NOT EDIT.

package tagging_service

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

// NewListAllTagsParams creates a new ListAllTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAllTagsParams() *ListAllTagsParams {
	return &ListAllTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAllTagsParamsWithTimeout creates a new ListAllTagsParams object
// with the ability to set a timeout on a request.
func NewListAllTagsParamsWithTimeout(timeout time.Duration) *ListAllTagsParams {
	return &ListAllTagsParams{
		timeout: timeout,
	}
}

// NewListAllTagsParamsWithContext creates a new ListAllTagsParams object
// with the ability to set a context for a request.
func NewListAllTagsParamsWithContext(ctx context.Context) *ListAllTagsParams {
	return &ListAllTagsParams{
		Context: ctx,
	}
}

// NewListAllTagsParamsWithHTTPClient creates a new ListAllTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAllTagsParamsWithHTTPClient(client *http.Client) *ListAllTagsParams {
	return &ListAllTagsParams{
		HTTPClient: client,
	}
}

/*
ListAllTagsParams contains all the parameters to send to the API endpoint

	for the list all tags operation.

	Typically these are written to a http.Request.
*/
type ListAllTagsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list all tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAllTagsParams) WithDefaults() *ListAllTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list all tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAllTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list all tags params
func (o *ListAllTagsParams) WithTimeout(timeout time.Duration) *ListAllTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all tags params
func (o *ListAllTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all tags params
func (o *ListAllTagsParams) WithContext(ctx context.Context) *ListAllTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all tags params
func (o *ListAllTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all tags params
func (o *ListAllTagsParams) WithHTTPClient(client *http.Client) *ListAllTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all tags params
func (o *ListAllTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
