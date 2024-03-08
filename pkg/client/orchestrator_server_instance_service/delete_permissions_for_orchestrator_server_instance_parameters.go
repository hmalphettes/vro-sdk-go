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

// NewDeletePermissionsForOrchestratorServerInstanceParams creates a new DeletePermissionsForOrchestratorServerInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePermissionsForOrchestratorServerInstanceParams() *DeletePermissionsForOrchestratorServerInstanceParams {
	return &DeletePermissionsForOrchestratorServerInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePermissionsForOrchestratorServerInstanceParamsWithTimeout creates a new DeletePermissionsForOrchestratorServerInstanceParams object
// with the ability to set a timeout on a request.
func NewDeletePermissionsForOrchestratorServerInstanceParamsWithTimeout(timeout time.Duration) *DeletePermissionsForOrchestratorServerInstanceParams {
	return &DeletePermissionsForOrchestratorServerInstanceParams{
		timeout: timeout,
	}
}

// NewDeletePermissionsForOrchestratorServerInstanceParamsWithContext creates a new DeletePermissionsForOrchestratorServerInstanceParams object
// with the ability to set a context for a request.
func NewDeletePermissionsForOrchestratorServerInstanceParamsWithContext(ctx context.Context) *DeletePermissionsForOrchestratorServerInstanceParams {
	return &DeletePermissionsForOrchestratorServerInstanceParams{
		Context: ctx,
	}
}

// NewDeletePermissionsForOrchestratorServerInstanceParamsWithHTTPClient creates a new DeletePermissionsForOrchestratorServerInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePermissionsForOrchestratorServerInstanceParamsWithHTTPClient(client *http.Client) *DeletePermissionsForOrchestratorServerInstanceParams {
	return &DeletePermissionsForOrchestratorServerInstanceParams{
		HTTPClient: client,
	}
}

/*
DeletePermissionsForOrchestratorServerInstanceParams contains all the parameters to send to the API endpoint

	for the delete permissions for orchestrator server instance operation.

	Typically these are written to a http.Request.
*/
type DeletePermissionsForOrchestratorServerInstanceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete permissions for orchestrator server instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePermissionsForOrchestratorServerInstanceParams) WithDefaults() *DeletePermissionsForOrchestratorServerInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete permissions for orchestrator server instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePermissionsForOrchestratorServerInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete permissions for orchestrator server instance params
func (o *DeletePermissionsForOrchestratorServerInstanceParams) WithTimeout(timeout time.Duration) *DeletePermissionsForOrchestratorServerInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete permissions for orchestrator server instance params
func (o *DeletePermissionsForOrchestratorServerInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete permissions for orchestrator server instance params
func (o *DeletePermissionsForOrchestratorServerInstanceParams) WithContext(ctx context.Context) *DeletePermissionsForOrchestratorServerInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete permissions for orchestrator server instance params
func (o *DeletePermissionsForOrchestratorServerInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete permissions for orchestrator server instance params
func (o *DeletePermissionsForOrchestratorServerInstanceParams) WithHTTPClient(client *http.Client) *DeletePermissionsForOrchestratorServerInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete permissions for orchestrator server instance params
func (o *DeletePermissionsForOrchestratorServerInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePermissionsForOrchestratorServerInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}