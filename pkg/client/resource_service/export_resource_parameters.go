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

// NewExportResourceParams creates a new ExportResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportResourceParams() *ExportResourceParams {
	return &ExportResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportResourceParamsWithTimeout creates a new ExportResourceParams object
// with the ability to set a timeout on a request.
func NewExportResourceParamsWithTimeout(timeout time.Duration) *ExportResourceParams {
	return &ExportResourceParams{
		timeout: timeout,
	}
}

// NewExportResourceParamsWithContext creates a new ExportResourceParams object
// with the ability to set a context for a request.
func NewExportResourceParamsWithContext(ctx context.Context) *ExportResourceParams {
	return &ExportResourceParams{
		Context: ctx,
	}
}

// NewExportResourceParamsWithHTTPClient creates a new ExportResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportResourceParamsWithHTTPClient(client *http.Client) *ExportResourceParams {
	return &ExportResourceParams{
		HTTPClient: client,
	}
}

/*
ExportResourceParams contains all the parameters to send to the API endpoint

	for the export resource operation.

	Typically these are written to a http.Request.
*/
type ExportResourceParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportResourceParams) WithDefaults() *ExportResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export resource params
func (o *ExportResourceParams) WithTimeout(timeout time.Duration) *ExportResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export resource params
func (o *ExportResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export resource params
func (o *ExportResourceParams) WithContext(ctx context.Context) *ExportResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export resource params
func (o *ExportResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export resource params
func (o *ExportResourceParams) WithHTTPClient(client *http.Client) *ExportResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export resource params
func (o *ExportResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the export resource params
func (o *ExportResourceParams) WithID(id string) *ExportResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the export resource params
func (o *ExportResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExportResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
