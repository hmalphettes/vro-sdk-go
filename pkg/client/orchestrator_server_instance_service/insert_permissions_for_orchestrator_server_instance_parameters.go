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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewInsertPermissionsForOrchestratorServerInstanceParams creates a new InsertPermissionsForOrchestratorServerInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInsertPermissionsForOrchestratorServerInstanceParams() *InsertPermissionsForOrchestratorServerInstanceParams {
	return &InsertPermissionsForOrchestratorServerInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInsertPermissionsForOrchestratorServerInstanceParamsWithTimeout creates a new InsertPermissionsForOrchestratorServerInstanceParams object
// with the ability to set a timeout on a request.
func NewInsertPermissionsForOrchestratorServerInstanceParamsWithTimeout(timeout time.Duration) *InsertPermissionsForOrchestratorServerInstanceParams {
	return &InsertPermissionsForOrchestratorServerInstanceParams{
		timeout: timeout,
	}
}

// NewInsertPermissionsForOrchestratorServerInstanceParamsWithContext creates a new InsertPermissionsForOrchestratorServerInstanceParams object
// with the ability to set a context for a request.
func NewInsertPermissionsForOrchestratorServerInstanceParamsWithContext(ctx context.Context) *InsertPermissionsForOrchestratorServerInstanceParams {
	return &InsertPermissionsForOrchestratorServerInstanceParams{
		Context: ctx,
	}
}

// NewInsertPermissionsForOrchestratorServerInstanceParamsWithHTTPClient creates a new InsertPermissionsForOrchestratorServerInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewInsertPermissionsForOrchestratorServerInstanceParamsWithHTTPClient(client *http.Client) *InsertPermissionsForOrchestratorServerInstanceParams {
	return &InsertPermissionsForOrchestratorServerInstanceParams{
		HTTPClient: client,
	}
}

/*
InsertPermissionsForOrchestratorServerInstanceParams contains all the parameters to send to the API endpoint

	for the insert permissions for orchestrator server instance operation.

	Typically these are written to a http.Request.
*/
type InsertPermissionsForOrchestratorServerInstanceParams struct {

	// Body.
	Body *models.Permissions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the insert permissions for orchestrator server instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertPermissionsForOrchestratorServerInstanceParams) WithDefaults() *InsertPermissionsForOrchestratorServerInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the insert permissions for orchestrator server instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertPermissionsForOrchestratorServerInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) WithTimeout(timeout time.Duration) *InsertPermissionsForOrchestratorServerInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) WithContext(ctx context.Context) *InsertPermissionsForOrchestratorServerInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) WithHTTPClient(client *http.Client) *InsertPermissionsForOrchestratorServerInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) WithBody(body *models.Permissions) *InsertPermissionsForOrchestratorServerInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the insert permissions for orchestrator server instance params
func (o *InsertPermissionsForOrchestratorServerInstanceParams) SetBody(body *models.Permissions) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InsertPermissionsForOrchestratorServerInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}