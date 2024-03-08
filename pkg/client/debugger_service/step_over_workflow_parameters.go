// Code generated by go-swagger; DO NOT EDIT.

package debugger_service

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

// NewStepOverWorkflowParams creates a new StepOverWorkflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStepOverWorkflowParams() *StepOverWorkflowParams {
	return &StepOverWorkflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStepOverWorkflowParamsWithTimeout creates a new StepOverWorkflowParams object
// with the ability to set a timeout on a request.
func NewStepOverWorkflowParamsWithTimeout(timeout time.Duration) *StepOverWorkflowParams {
	return &StepOverWorkflowParams{
		timeout: timeout,
	}
}

// NewStepOverWorkflowParamsWithContext creates a new StepOverWorkflowParams object
// with the ability to set a context for a request.
func NewStepOverWorkflowParamsWithContext(ctx context.Context) *StepOverWorkflowParams {
	return &StepOverWorkflowParams{
		Context: ctx,
	}
}

// NewStepOverWorkflowParamsWithHTTPClient creates a new StepOverWorkflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewStepOverWorkflowParamsWithHTTPClient(client *http.Client) *StepOverWorkflowParams {
	return &StepOverWorkflowParams{
		HTTPClient: client,
	}
}

/*
StepOverWorkflowParams contains all the parameters to send to the API endpoint

	for the step over workflow operation.

	Typically these are written to a http.Request.
*/
type StepOverWorkflowParams struct {

	// ExecutionID.
	ExecutionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the step over workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StepOverWorkflowParams) WithDefaults() *StepOverWorkflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the step over workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StepOverWorkflowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the step over workflow params
func (o *StepOverWorkflowParams) WithTimeout(timeout time.Duration) *StepOverWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the step over workflow params
func (o *StepOverWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the step over workflow params
func (o *StepOverWorkflowParams) WithContext(ctx context.Context) *StepOverWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the step over workflow params
func (o *StepOverWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the step over workflow params
func (o *StepOverWorkflowParams) WithHTTPClient(client *http.Client) *StepOverWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the step over workflow params
func (o *StepOverWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the step over workflow params
func (o *StepOverWorkflowParams) WithExecutionID(executionID string) *StepOverWorkflowParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the step over workflow params
func (o *StepOverWorkflowParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WriteToRequest writes these params to a swagger request
func (o *StepOverWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionId
	if err := r.SetPathParam("executionId", o.ExecutionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
