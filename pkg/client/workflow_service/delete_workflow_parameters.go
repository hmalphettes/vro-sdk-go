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
	"github.com/go-openapi/swag"
)

// NewDeleteWorkflowParams creates a new DeleteWorkflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWorkflowParams() *DeleteWorkflowParams {
	return &DeleteWorkflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkflowParamsWithTimeout creates a new DeleteWorkflowParams object
// with the ability to set a timeout on a request.
func NewDeleteWorkflowParamsWithTimeout(timeout time.Duration) *DeleteWorkflowParams {
	return &DeleteWorkflowParams{
		timeout: timeout,
	}
}

// NewDeleteWorkflowParamsWithContext creates a new DeleteWorkflowParams object
// with the ability to set a context for a request.
func NewDeleteWorkflowParamsWithContext(ctx context.Context) *DeleteWorkflowParams {
	return &DeleteWorkflowParams{
		Context: ctx,
	}
}

// NewDeleteWorkflowParamsWithHTTPClient creates a new DeleteWorkflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWorkflowParamsWithHTTPClient(client *http.Client) *DeleteWorkflowParams {
	return &DeleteWorkflowParams{
		HTTPClient: client,
	}
}

/*
DeleteWorkflowParams contains all the parameters to send to the API endpoint

	for the delete workflow operation.

	Typically these are written to a http.Request.
*/
type DeleteWorkflowParams struct {

	// Force.
	Force *bool

	// ForceDeleteLocked.
	ForceDeleteLocked *bool

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkflowParams) WithDefaults() *DeleteWorkflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkflowParams) SetDefaults() {
	var (
		forceDefault = bool(false)

		forceDeleteLockedDefault = bool(false)
	)

	val := DeleteWorkflowParams{
		Force:             &forceDefault,
		ForceDeleteLocked: &forceDeleteLockedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete workflow params
func (o *DeleteWorkflowParams) WithTimeout(timeout time.Duration) *DeleteWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workflow params
func (o *DeleteWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workflow params
func (o *DeleteWorkflowParams) WithContext(ctx context.Context) *DeleteWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workflow params
func (o *DeleteWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workflow params
func (o *DeleteWorkflowParams) WithHTTPClient(client *http.Client) *DeleteWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workflow params
func (o *DeleteWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the delete workflow params
func (o *DeleteWorkflowParams) WithForce(force *bool) *DeleteWorkflowParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete workflow params
func (o *DeleteWorkflowParams) SetForce(force *bool) {
	o.Force = force
}

// WithForceDeleteLocked adds the forceDeleteLocked to the delete workflow params
func (o *DeleteWorkflowParams) WithForceDeleteLocked(forceDeleteLocked *bool) *DeleteWorkflowParams {
	o.SetForceDeleteLocked(forceDeleteLocked)
	return o
}

// SetForceDeleteLocked adds the forceDeleteLocked to the delete workflow params
func (o *DeleteWorkflowParams) SetForceDeleteLocked(forceDeleteLocked *bool) {
	o.ForceDeleteLocked = forceDeleteLocked
}

// WithWorkflowID adds the workflowID to the delete workflow params
func (o *DeleteWorkflowParams) WithWorkflowID(workflowID string) *DeleteWorkflowParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the delete workflow params
func (o *DeleteWorkflowParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	if o.ForceDeleteLocked != nil {

		// query param forceDeleteLocked
		var qrForceDeleteLocked bool

		if o.ForceDeleteLocked != nil {
			qrForceDeleteLocked = *o.ForceDeleteLocked
		}
		qForceDeleteLocked := swag.FormatBool(qrForceDeleteLocked)
		if qForceDeleteLocked != "" {

			if err := r.SetQueryParam("forceDeleteLocked", qForceDeleteLocked); err != nil {
				return err
			}
		}
	}

	// path param workflowId
	if err := r.SetPathParam("workflowId", o.WorkflowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}