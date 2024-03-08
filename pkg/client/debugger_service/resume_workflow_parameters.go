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

// NewResumeWorkflowParams creates a new ResumeWorkflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResumeWorkflowParams() *ResumeWorkflowParams {
	return &ResumeWorkflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResumeWorkflowParamsWithTimeout creates a new ResumeWorkflowParams object
// with the ability to set a timeout on a request.
func NewResumeWorkflowParamsWithTimeout(timeout time.Duration) *ResumeWorkflowParams {
	return &ResumeWorkflowParams{
		timeout: timeout,
	}
}

// NewResumeWorkflowParamsWithContext creates a new ResumeWorkflowParams object
// with the ability to set a context for a request.
func NewResumeWorkflowParamsWithContext(ctx context.Context) *ResumeWorkflowParams {
	return &ResumeWorkflowParams{
		Context: ctx,
	}
}

// NewResumeWorkflowParamsWithHTTPClient creates a new ResumeWorkflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewResumeWorkflowParamsWithHTTPClient(client *http.Client) *ResumeWorkflowParams {
	return &ResumeWorkflowParams{
		HTTPClient: client,
	}
}

/*
ResumeWorkflowParams contains all the parameters to send to the API endpoint

	for the resume workflow operation.

	Typically these are written to a http.Request.
*/
type ResumeWorkflowParams struct {

	// ExecutionID.
	ExecutionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resume workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeWorkflowParams) WithDefaults() *ResumeWorkflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resume workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeWorkflowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resume workflow params
func (o *ResumeWorkflowParams) WithTimeout(timeout time.Duration) *ResumeWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume workflow params
func (o *ResumeWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume workflow params
func (o *ResumeWorkflowParams) WithContext(ctx context.Context) *ResumeWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume workflow params
func (o *ResumeWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume workflow params
func (o *ResumeWorkflowParams) WithHTTPClient(client *http.Client) *ResumeWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume workflow params
func (o *ResumeWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the resume workflow params
func (o *ResumeWorkflowParams) WithExecutionID(executionID string) *ResumeWorkflowParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the resume workflow params
func (o *ResumeWorkflowParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
