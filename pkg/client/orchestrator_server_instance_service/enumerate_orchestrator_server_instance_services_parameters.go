// Code generated by go-swagger; DO NOT EDIT.

package orchestrator_server_instance_service

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

// NewEnumerateOrchestratorServerInstanceServicesParams creates a new EnumerateOrchestratorServerInstanceServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnumerateOrchestratorServerInstanceServicesParams() *EnumerateOrchestratorServerInstanceServicesParams {
	return &EnumerateOrchestratorServerInstanceServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnumerateOrchestratorServerInstanceServicesParamsWithTimeout creates a new EnumerateOrchestratorServerInstanceServicesParams object
// with the ability to set a timeout on a request.
func NewEnumerateOrchestratorServerInstanceServicesParamsWithTimeout(timeout time.Duration) *EnumerateOrchestratorServerInstanceServicesParams {
	return &EnumerateOrchestratorServerInstanceServicesParams{
		timeout: timeout,
	}
}

// NewEnumerateOrchestratorServerInstanceServicesParamsWithContext creates a new EnumerateOrchestratorServerInstanceServicesParams object
// with the ability to set a context for a request.
func NewEnumerateOrchestratorServerInstanceServicesParamsWithContext(ctx context.Context) *EnumerateOrchestratorServerInstanceServicesParams {
	return &EnumerateOrchestratorServerInstanceServicesParams{
		Context: ctx,
	}
}

// NewEnumerateOrchestratorServerInstanceServicesParamsWithHTTPClient creates a new EnumerateOrchestratorServerInstanceServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnumerateOrchestratorServerInstanceServicesParamsWithHTTPClient(client *http.Client) *EnumerateOrchestratorServerInstanceServicesParams {
	return &EnumerateOrchestratorServerInstanceServicesParams{
		HTTPClient: client,
	}
}

/*
EnumerateOrchestratorServerInstanceServicesParams contains all the parameters to send to the API endpoint

	for the enumerate orchestrator server instance services operation.

	Typically these are written to a http.Request.
*/
type EnumerateOrchestratorServerInstanceServicesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enumerate orchestrator server instance services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumerateOrchestratorServerInstanceServicesParams) WithDefaults() *EnumerateOrchestratorServerInstanceServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enumerate orchestrator server instance services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumerateOrchestratorServerInstanceServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enumerate orchestrator server instance services params
func (o *EnumerateOrchestratorServerInstanceServicesParams) WithTimeout(timeout time.Duration) *EnumerateOrchestratorServerInstanceServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate orchestrator server instance services params
func (o *EnumerateOrchestratorServerInstanceServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate orchestrator server instance services params
func (o *EnumerateOrchestratorServerInstanceServicesParams) WithContext(ctx context.Context) *EnumerateOrchestratorServerInstanceServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate orchestrator server instance services params
func (o *EnumerateOrchestratorServerInstanceServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate orchestrator server instance services params
func (o *EnumerateOrchestratorServerInstanceServicesParams) WithHTTPClient(client *http.Client) *EnumerateOrchestratorServerInstanceServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate orchestrator server instance services params
func (o *EnumerateOrchestratorServerInstanceServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EnumerateOrchestratorServerInstanceServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}