// Code generated by go-swagger; DO NOT EDIT.

package workflow_service

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

// NewDownloadWorkflowSchemaContentParams creates a new DownloadWorkflowSchemaContentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadWorkflowSchemaContentParams() *DownloadWorkflowSchemaContentParams {
	return &DownloadWorkflowSchemaContentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadWorkflowSchemaContentParamsWithTimeout creates a new DownloadWorkflowSchemaContentParams object
// with the ability to set a timeout on a request.
func NewDownloadWorkflowSchemaContentParamsWithTimeout(timeout time.Duration) *DownloadWorkflowSchemaContentParams {
	return &DownloadWorkflowSchemaContentParams{
		timeout: timeout,
	}
}

// NewDownloadWorkflowSchemaContentParamsWithContext creates a new DownloadWorkflowSchemaContentParams object
// with the ability to set a context for a request.
func NewDownloadWorkflowSchemaContentParamsWithContext(ctx context.Context) *DownloadWorkflowSchemaContentParams {
	return &DownloadWorkflowSchemaContentParams{
		Context: ctx,
	}
}

// NewDownloadWorkflowSchemaContentParamsWithHTTPClient creates a new DownloadWorkflowSchemaContentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadWorkflowSchemaContentParamsWithHTTPClient(client *http.Client) *DownloadWorkflowSchemaContentParams {
	return &DownloadWorkflowSchemaContentParams{
		HTTPClient: client,
	}
}

/*
DownloadWorkflowSchemaContentParams contains all the parameters to send to the API endpoint

	for the download workflow schema content operation.

	Typically these are written to a http.Request.
*/
type DownloadWorkflowSchemaContentParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download workflow schema content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadWorkflowSchemaContentParams) WithDefaults() *DownloadWorkflowSchemaContentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download workflow schema content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadWorkflowSchemaContentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) WithTimeout(timeout time.Duration) *DownloadWorkflowSchemaContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) WithContext(ctx context.Context) *DownloadWorkflowSchemaContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) WithHTTPClient(client *http.Client) *DownloadWorkflowSchemaContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) WithID(id string) *DownloadWorkflowSchemaContentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download workflow schema content params
func (o *DownloadWorkflowSchemaContentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadWorkflowSchemaContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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