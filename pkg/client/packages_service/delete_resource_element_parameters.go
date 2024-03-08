// Code generated by go-swagger; DO NOT EDIT.

package packages_service

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

// NewDeleteResourceElementParams creates a new DeleteResourceElementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteResourceElementParams() *DeleteResourceElementParams {
	return &DeleteResourceElementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourceElementParamsWithTimeout creates a new DeleteResourceElementParams object
// with the ability to set a timeout on a request.
func NewDeleteResourceElementParamsWithTimeout(timeout time.Duration) *DeleteResourceElementParams {
	return &DeleteResourceElementParams{
		timeout: timeout,
	}
}

// NewDeleteResourceElementParamsWithContext creates a new DeleteResourceElementParams object
// with the ability to set a context for a request.
func NewDeleteResourceElementParamsWithContext(ctx context.Context) *DeleteResourceElementParams {
	return &DeleteResourceElementParams{
		Context: ctx,
	}
}

// NewDeleteResourceElementParamsWithHTTPClient creates a new DeleteResourceElementParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteResourceElementParamsWithHTTPClient(client *http.Client) *DeleteResourceElementParams {
	return &DeleteResourceElementParams{
		HTTPClient: client,
	}
}

/*
DeleteResourceElementParams contains all the parameters to send to the API endpoint

	for the delete resource element operation.

	Typically these are written to a http.Request.
*/
type DeleteResourceElementParams struct {

	// ID.
	ID string

	// PackageName.
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete resource element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceElementParams) WithDefaults() *DeleteResourceElementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete resource element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceElementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete resource element params
func (o *DeleteResourceElementParams) WithTimeout(timeout time.Duration) *DeleteResourceElementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete resource element params
func (o *DeleteResourceElementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete resource element params
func (o *DeleteResourceElementParams) WithContext(ctx context.Context) *DeleteResourceElementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete resource element params
func (o *DeleteResourceElementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete resource element params
func (o *DeleteResourceElementParams) WithHTTPClient(client *http.Client) *DeleteResourceElementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete resource element params
func (o *DeleteResourceElementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete resource element params
func (o *DeleteResourceElementParams) WithID(id string) *DeleteResourceElementParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete resource element params
func (o *DeleteResourceElementParams) SetID(id string) {
	o.ID = id
}

// WithPackageName adds the packageName to the delete resource element params
func (o *DeleteResourceElementParams) WithPackageName(packageName string) *DeleteResourceElementParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the delete resource element params
func (o *DeleteResourceElementParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourceElementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}