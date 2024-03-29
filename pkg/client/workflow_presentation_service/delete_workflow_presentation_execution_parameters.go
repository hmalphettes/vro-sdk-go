// Code generated by go-swagger; DO NOT EDIT.

package workflow_presentation_service

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

// NewDeleteWorkflowPresentationExecutionParams creates a new DeleteWorkflowPresentationExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWorkflowPresentationExecutionParams() *DeleteWorkflowPresentationExecutionParams {
	return &DeleteWorkflowPresentationExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkflowPresentationExecutionParamsWithTimeout creates a new DeleteWorkflowPresentationExecutionParams object
// with the ability to set a timeout on a request.
func NewDeleteWorkflowPresentationExecutionParamsWithTimeout(timeout time.Duration) *DeleteWorkflowPresentationExecutionParams {
	return &DeleteWorkflowPresentationExecutionParams{
		timeout: timeout,
	}
}

// NewDeleteWorkflowPresentationExecutionParamsWithContext creates a new DeleteWorkflowPresentationExecutionParams object
// with the ability to set a context for a request.
func NewDeleteWorkflowPresentationExecutionParamsWithContext(ctx context.Context) *DeleteWorkflowPresentationExecutionParams {
	return &DeleteWorkflowPresentationExecutionParams{
		Context: ctx,
	}
}

// NewDeleteWorkflowPresentationExecutionParamsWithHTTPClient creates a new DeleteWorkflowPresentationExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWorkflowPresentationExecutionParamsWithHTTPClient(client *http.Client) *DeleteWorkflowPresentationExecutionParams {
	return &DeleteWorkflowPresentationExecutionParams{
		HTTPClient: client,
	}
}

/*
DeleteWorkflowPresentationExecutionParams contains all the parameters to send to the API endpoint

	for the delete workflow presentation execution operation.

	Typically these are written to a http.Request.
*/
type DeleteWorkflowPresentationExecutionParams struct {

	// ExecutionID.
	ExecutionID string

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete workflow presentation execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkflowPresentationExecutionParams) WithDefaults() *DeleteWorkflowPresentationExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete workflow presentation execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkflowPresentationExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) WithTimeout(timeout time.Duration) *DeleteWorkflowPresentationExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) WithContext(ctx context.Context) *DeleteWorkflowPresentationExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) WithHTTPClient(client *http.Client) *DeleteWorkflowPresentationExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) WithExecutionID(executionID string) *DeleteWorkflowPresentationExecutionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithWorkflowID adds the workflowID to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) WithWorkflowID(workflowID string) *DeleteWorkflowPresentationExecutionParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the delete workflow presentation execution params
func (o *DeleteWorkflowPresentationExecutionParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkflowPresentationExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionId
	if err := r.SetPathParam("executionId", o.ExecutionID); err != nil {
		return err
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
