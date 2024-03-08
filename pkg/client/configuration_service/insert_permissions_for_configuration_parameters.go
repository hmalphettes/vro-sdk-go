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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewInsertPermissionsForConfigurationParams creates a new InsertPermissionsForConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInsertPermissionsForConfigurationParams() *InsertPermissionsForConfigurationParams {
	return &InsertPermissionsForConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInsertPermissionsForConfigurationParamsWithTimeout creates a new InsertPermissionsForConfigurationParams object
// with the ability to set a timeout on a request.
func NewInsertPermissionsForConfigurationParamsWithTimeout(timeout time.Duration) *InsertPermissionsForConfigurationParams {
	return &InsertPermissionsForConfigurationParams{
		timeout: timeout,
	}
}

// NewInsertPermissionsForConfigurationParamsWithContext creates a new InsertPermissionsForConfigurationParams object
// with the ability to set a context for a request.
func NewInsertPermissionsForConfigurationParamsWithContext(ctx context.Context) *InsertPermissionsForConfigurationParams {
	return &InsertPermissionsForConfigurationParams{
		Context: ctx,
	}
}

// NewInsertPermissionsForConfigurationParamsWithHTTPClient creates a new InsertPermissionsForConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewInsertPermissionsForConfigurationParamsWithHTTPClient(client *http.Client) *InsertPermissionsForConfigurationParams {
	return &InsertPermissionsForConfigurationParams{
		HTTPClient: client,
	}
}

/*
InsertPermissionsForConfigurationParams contains all the parameters to send to the API endpoint

	for the insert permissions for configuration operation.

	Typically these are written to a http.Request.
*/
type InsertPermissionsForConfigurationParams struct {

	// Body.
	Body *models.Permissions

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the insert permissions for configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertPermissionsForConfigurationParams) WithDefaults() *InsertPermissionsForConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the insert permissions for configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertPermissionsForConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) WithTimeout(timeout time.Duration) *InsertPermissionsForConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) WithContext(ctx context.Context) *InsertPermissionsForConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) WithHTTPClient(client *http.Client) *InsertPermissionsForConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) WithBody(body *models.Permissions) *InsertPermissionsForConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) SetBody(body *models.Permissions) {
	o.Body = body
}

// WithID adds the id to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) WithID(id string) *InsertPermissionsForConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the insert permissions for configuration params
func (o *InsertPermissionsForConfigurationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InsertPermissionsForConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}