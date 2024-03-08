// Code generated by go-swagger; DO NOT EDIT.

package workflow_run_service

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

// NewGetWorkflowExecutionStatisticsParams creates a new GetWorkflowExecutionStatisticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowExecutionStatisticsParams() *GetWorkflowExecutionStatisticsParams {
	return &GetWorkflowExecutionStatisticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowExecutionStatisticsParamsWithTimeout creates a new GetWorkflowExecutionStatisticsParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowExecutionStatisticsParamsWithTimeout(timeout time.Duration) *GetWorkflowExecutionStatisticsParams {
	return &GetWorkflowExecutionStatisticsParams{
		timeout: timeout,
	}
}

// NewGetWorkflowExecutionStatisticsParamsWithContext creates a new GetWorkflowExecutionStatisticsParams object
// with the ability to set a context for a request.
func NewGetWorkflowExecutionStatisticsParamsWithContext(ctx context.Context) *GetWorkflowExecutionStatisticsParams {
	return &GetWorkflowExecutionStatisticsParams{
		Context: ctx,
	}
}

// NewGetWorkflowExecutionStatisticsParamsWithHTTPClient creates a new GetWorkflowExecutionStatisticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowExecutionStatisticsParamsWithHTTPClient(client *http.Client) *GetWorkflowExecutionStatisticsParams {
	return &GetWorkflowExecutionStatisticsParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowExecutionStatisticsParams contains all the parameters to send to the API endpoint

	for the get workflow execution statistics operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowExecutionStatisticsParams struct {

	// ExecutionID.
	ExecutionID string

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow execution statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowExecutionStatisticsParams) WithDefaults() *GetWorkflowExecutionStatisticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow execution statistics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowExecutionStatisticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) WithTimeout(timeout time.Duration) *GetWorkflowExecutionStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) WithContext(ctx context.Context) *GetWorkflowExecutionStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) WithHTTPClient(client *http.Client) *GetWorkflowExecutionStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) WithExecutionID(executionID string) *GetWorkflowExecutionStatisticsParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithWorkflowID adds the workflowID to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) WithWorkflowID(workflowID string) *GetWorkflowExecutionStatisticsParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the get workflow execution statistics params
func (o *GetWorkflowExecutionStatisticsParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowExecutionStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
