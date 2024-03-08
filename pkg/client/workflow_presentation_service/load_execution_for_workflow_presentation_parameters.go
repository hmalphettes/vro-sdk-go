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

// NewLoadExecutionForWorkflowPresentationParams creates a new LoadExecutionForWorkflowPresentationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLoadExecutionForWorkflowPresentationParams() *LoadExecutionForWorkflowPresentationParams {
	return &LoadExecutionForWorkflowPresentationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLoadExecutionForWorkflowPresentationParamsWithTimeout creates a new LoadExecutionForWorkflowPresentationParams object
// with the ability to set a timeout on a request.
func NewLoadExecutionForWorkflowPresentationParamsWithTimeout(timeout time.Duration) *LoadExecutionForWorkflowPresentationParams {
	return &LoadExecutionForWorkflowPresentationParams{
		timeout: timeout,
	}
}

// NewLoadExecutionForWorkflowPresentationParamsWithContext creates a new LoadExecutionForWorkflowPresentationParams object
// with the ability to set a context for a request.
func NewLoadExecutionForWorkflowPresentationParamsWithContext(ctx context.Context) *LoadExecutionForWorkflowPresentationParams {
	return &LoadExecutionForWorkflowPresentationParams{
		Context: ctx,
	}
}

// NewLoadExecutionForWorkflowPresentationParamsWithHTTPClient creates a new LoadExecutionForWorkflowPresentationParams object
// with the ability to set a custom HTTPClient for a request.
func NewLoadExecutionForWorkflowPresentationParamsWithHTTPClient(client *http.Client) *LoadExecutionForWorkflowPresentationParams {
	return &LoadExecutionForWorkflowPresentationParams{
		HTTPClient: client,
	}
}

/*
LoadExecutionForWorkflowPresentationParams contains all the parameters to send to the API endpoint

	for the load execution for workflow presentation operation.

	Typically these are written to a http.Request.
*/
type LoadExecutionForWorkflowPresentationParams struct {

	// ExecutionID.
	ExecutionID string

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the load execution for workflow presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoadExecutionForWorkflowPresentationParams) WithDefaults() *LoadExecutionForWorkflowPresentationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the load execution for workflow presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoadExecutionForWorkflowPresentationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) WithTimeout(timeout time.Duration) *LoadExecutionForWorkflowPresentationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) WithContext(ctx context.Context) *LoadExecutionForWorkflowPresentationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) WithHTTPClient(client *http.Client) *LoadExecutionForWorkflowPresentationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) WithExecutionID(executionID string) *LoadExecutionForWorkflowPresentationParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithWorkflowID adds the workflowID to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) WithWorkflowID(workflowID string) *LoadExecutionForWorkflowPresentationParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the load execution for workflow presentation params
func (o *LoadExecutionForWorkflowPresentationParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *LoadExecutionForWorkflowPresentationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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