// Code generated by go-swagger; DO NOT EDIT.

package actions_service

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

// NewExportActionParams creates a new ExportActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportActionParams() *ExportActionParams {
	return &ExportActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportActionParamsWithTimeout creates a new ExportActionParams object
// with the ability to set a timeout on a request.
func NewExportActionParamsWithTimeout(timeout time.Duration) *ExportActionParams {
	return &ExportActionParams{
		timeout: timeout,
	}
}

// NewExportActionParamsWithContext creates a new ExportActionParams object
// with the ability to set a context for a request.
func NewExportActionParamsWithContext(ctx context.Context) *ExportActionParams {
	return &ExportActionParams{
		Context: ctx,
	}
}

// NewExportActionParamsWithHTTPClient creates a new ExportActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportActionParamsWithHTTPClient(client *http.Client) *ExportActionParams {
	return &ExportActionParams{
		HTTPClient: client,
	}
}

/*
ExportActionParams contains all the parameters to send to the API endpoint

	for the export action operation.

	Typically these are written to a http.Request.
*/
type ExportActionParams struct {

	// ActionID.
	ActionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportActionParams) WithDefaults() *ExportActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export action params
func (o *ExportActionParams) WithTimeout(timeout time.Duration) *ExportActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export action params
func (o *ExportActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export action params
func (o *ExportActionParams) WithContext(ctx context.Context) *ExportActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export action params
func (o *ExportActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export action params
func (o *ExportActionParams) WithHTTPClient(client *http.Client) *ExportActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export action params
func (o *ExportActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the export action params
func (o *ExportActionParams) WithActionID(actionID string) *ExportActionParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the export action params
func (o *ExportActionParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
