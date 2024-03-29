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

// NewDeletePermissionsForResourceParams creates a new DeletePermissionsForResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePermissionsForResourceParams() *DeletePermissionsForResourceParams {
	return &DeletePermissionsForResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePermissionsForResourceParamsWithTimeout creates a new DeletePermissionsForResourceParams object
// with the ability to set a timeout on a request.
func NewDeletePermissionsForResourceParamsWithTimeout(timeout time.Duration) *DeletePermissionsForResourceParams {
	return &DeletePermissionsForResourceParams{
		timeout: timeout,
	}
}

// NewDeletePermissionsForResourceParamsWithContext creates a new DeletePermissionsForResourceParams object
// with the ability to set a context for a request.
func NewDeletePermissionsForResourceParamsWithContext(ctx context.Context) *DeletePermissionsForResourceParams {
	return &DeletePermissionsForResourceParams{
		Context: ctx,
	}
}

// NewDeletePermissionsForResourceParamsWithHTTPClient creates a new DeletePermissionsForResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePermissionsForResourceParamsWithHTTPClient(client *http.Client) *DeletePermissionsForResourceParams {
	return &DeletePermissionsForResourceParams{
		HTTPClient: client,
	}
}

/*
DeletePermissionsForResourceParams contains all the parameters to send to the API endpoint

	for the delete permissions for resource operation.

	Typically these are written to a http.Request.
*/
type DeletePermissionsForResourceParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete permissions for resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePermissionsForResourceParams) WithDefaults() *DeletePermissionsForResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete permissions for resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePermissionsForResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) WithTimeout(timeout time.Duration) *DeletePermissionsForResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) WithContext(ctx context.Context) *DeletePermissionsForResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) WithHTTPClient(client *http.Client) *DeletePermissionsForResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) WithID(id string) *DeletePermissionsForResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete permissions for resource params
func (o *DeletePermissionsForResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePermissionsForResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
