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

// NewDownloadWorkflowIconParams creates a new DownloadWorkflowIconParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadWorkflowIconParams() *DownloadWorkflowIconParams {
	return &DownloadWorkflowIconParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadWorkflowIconParamsWithTimeout creates a new DownloadWorkflowIconParams object
// with the ability to set a timeout on a request.
func NewDownloadWorkflowIconParamsWithTimeout(timeout time.Duration) *DownloadWorkflowIconParams {
	return &DownloadWorkflowIconParams{
		timeout: timeout,
	}
}

// NewDownloadWorkflowIconParamsWithContext creates a new DownloadWorkflowIconParams object
// with the ability to set a context for a request.
func NewDownloadWorkflowIconParamsWithContext(ctx context.Context) *DownloadWorkflowIconParams {
	return &DownloadWorkflowIconParams{
		Context: ctx,
	}
}

// NewDownloadWorkflowIconParamsWithHTTPClient creates a new DownloadWorkflowIconParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadWorkflowIconParamsWithHTTPClient(client *http.Client) *DownloadWorkflowIconParams {
	return &DownloadWorkflowIconParams{
		HTTPClient: client,
	}
}

/*
DownloadWorkflowIconParams contains all the parameters to send to the API endpoint

	for the download workflow icon operation.

	Typically these are written to a http.Request.
*/
type DownloadWorkflowIconParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download workflow icon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadWorkflowIconParams) WithDefaults() *DownloadWorkflowIconParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download workflow icon params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadWorkflowIconParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download workflow icon params
func (o *DownloadWorkflowIconParams) WithTimeout(timeout time.Duration) *DownloadWorkflowIconParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download workflow icon params
func (o *DownloadWorkflowIconParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download workflow icon params
func (o *DownloadWorkflowIconParams) WithContext(ctx context.Context) *DownloadWorkflowIconParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download workflow icon params
func (o *DownloadWorkflowIconParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download workflow icon params
func (o *DownloadWorkflowIconParams) WithHTTPClient(client *http.Client) *DownloadWorkflowIconParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download workflow icon params
func (o *DownloadWorkflowIconParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the download workflow icon params
func (o *DownloadWorkflowIconParams) WithID(id string) *DownloadWorkflowIconParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download workflow icon params
func (o *DownloadWorkflowIconParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadWorkflowIconParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
