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

// NewDeleteActionElementParams creates a new DeleteActionElementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteActionElementParams() *DeleteActionElementParams {
	return &DeleteActionElementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteActionElementParamsWithTimeout creates a new DeleteActionElementParams object
// with the ability to set a timeout on a request.
func NewDeleteActionElementParamsWithTimeout(timeout time.Duration) *DeleteActionElementParams {
	return &DeleteActionElementParams{
		timeout: timeout,
	}
}

// NewDeleteActionElementParamsWithContext creates a new DeleteActionElementParams object
// with the ability to set a context for a request.
func NewDeleteActionElementParamsWithContext(ctx context.Context) *DeleteActionElementParams {
	return &DeleteActionElementParams{
		Context: ctx,
	}
}

// NewDeleteActionElementParamsWithHTTPClient creates a new DeleteActionElementParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteActionElementParamsWithHTTPClient(client *http.Client) *DeleteActionElementParams {
	return &DeleteActionElementParams{
		HTTPClient: client,
	}
}

/*
DeleteActionElementParams contains all the parameters to send to the API endpoint

	for the delete action element operation.

	Typically these are written to a http.Request.
*/
type DeleteActionElementParams struct {

	// ID.
	ID string

	// PackageName.
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete action element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteActionElementParams) WithDefaults() *DeleteActionElementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete action element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteActionElementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete action element params
func (o *DeleteActionElementParams) WithTimeout(timeout time.Duration) *DeleteActionElementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete action element params
func (o *DeleteActionElementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete action element params
func (o *DeleteActionElementParams) WithContext(ctx context.Context) *DeleteActionElementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete action element params
func (o *DeleteActionElementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete action element params
func (o *DeleteActionElementParams) WithHTTPClient(client *http.Client) *DeleteActionElementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete action element params
func (o *DeleteActionElementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete action element params
func (o *DeleteActionElementParams) WithID(id string) *DeleteActionElementParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete action element params
func (o *DeleteActionElementParams) SetID(id string) {
	o.ID = id
}

// WithPackageName adds the packageName to the delete action element params
func (o *DeleteActionElementParams) WithPackageName(packageName string) *DeleteActionElementParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the delete action element params
func (o *DeleteActionElementParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteActionElementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
