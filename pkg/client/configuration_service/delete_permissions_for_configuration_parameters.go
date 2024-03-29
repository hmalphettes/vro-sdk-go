// Code generated by go-swagger; DO NOT EDIT.

package configuration_service

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

// NewDeletePermissionsForConfigurationParams creates a new DeletePermissionsForConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePermissionsForConfigurationParams() *DeletePermissionsForConfigurationParams {
	return &DeletePermissionsForConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePermissionsForConfigurationParamsWithTimeout creates a new DeletePermissionsForConfigurationParams object
// with the ability to set a timeout on a request.
func NewDeletePermissionsForConfigurationParamsWithTimeout(timeout time.Duration) *DeletePermissionsForConfigurationParams {
	return &DeletePermissionsForConfigurationParams{
		timeout: timeout,
	}
}

// NewDeletePermissionsForConfigurationParamsWithContext creates a new DeletePermissionsForConfigurationParams object
// with the ability to set a context for a request.
func NewDeletePermissionsForConfigurationParamsWithContext(ctx context.Context) *DeletePermissionsForConfigurationParams {
	return &DeletePermissionsForConfigurationParams{
		Context: ctx,
	}
}

// NewDeletePermissionsForConfigurationParamsWithHTTPClient creates a new DeletePermissionsForConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePermissionsForConfigurationParamsWithHTTPClient(client *http.Client) *DeletePermissionsForConfigurationParams {
	return &DeletePermissionsForConfigurationParams{
		HTTPClient: client,
	}
}

/*
DeletePermissionsForConfigurationParams contains all the parameters to send to the API endpoint

	for the delete permissions for configuration operation.

	Typically these are written to a http.Request.
*/
type DeletePermissionsForConfigurationParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete permissions for configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePermissionsForConfigurationParams) WithDefaults() *DeletePermissionsForConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete permissions for configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePermissionsForConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) WithTimeout(timeout time.Duration) *DeletePermissionsForConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) WithContext(ctx context.Context) *DeletePermissionsForConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) WithHTTPClient(client *http.Client) *DeletePermissionsForConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) WithID(id string) *DeletePermissionsForConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete permissions for configuration params
func (o *DeletePermissionsForConfigurationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePermissionsForConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
