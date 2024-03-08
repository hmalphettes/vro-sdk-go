// Code generated by go-swagger; DO NOT EDIT.

package workflow_user_interaction_presentation_service

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

// NewGetPresentationForInteractionParams creates a new GetPresentationForInteractionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPresentationForInteractionParams() *GetPresentationForInteractionParams {
	return &GetPresentationForInteractionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPresentationForInteractionParamsWithTimeout creates a new GetPresentationForInteractionParams object
// with the ability to set a timeout on a request.
func NewGetPresentationForInteractionParamsWithTimeout(timeout time.Duration) *GetPresentationForInteractionParams {
	return &GetPresentationForInteractionParams{
		timeout: timeout,
	}
}

// NewGetPresentationForInteractionParamsWithContext creates a new GetPresentationForInteractionParams object
// with the ability to set a context for a request.
func NewGetPresentationForInteractionParamsWithContext(ctx context.Context) *GetPresentationForInteractionParams {
	return &GetPresentationForInteractionParams{
		Context: ctx,
	}
}

// NewGetPresentationForInteractionParamsWithHTTPClient creates a new GetPresentationForInteractionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPresentationForInteractionParamsWithHTTPClient(client *http.Client) *GetPresentationForInteractionParams {
	return &GetPresentationForInteractionParams{
		HTTPClient: client,
	}
}

/*
GetPresentationForInteractionParams contains all the parameters to send to the API endpoint

	for the get presentation for interaction operation.

	Typically these are written to a http.Request.
*/
type GetPresentationForInteractionParams struct {

	// ExecutionID.
	ExecutionID string

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get presentation for interaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPresentationForInteractionParams) WithDefaults() *GetPresentationForInteractionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get presentation for interaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPresentationForInteractionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) WithTimeout(timeout time.Duration) *GetPresentationForInteractionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) WithContext(ctx context.Context) *GetPresentationForInteractionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) WithHTTPClient(client *http.Client) *GetPresentationForInteractionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) WithExecutionID(executionID string) *GetPresentationForInteractionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithWorkflowID adds the workflowID to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) WithWorkflowID(workflowID string) *GetPresentationForInteractionParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the get presentation for interaction params
func (o *GetPresentationForInteractionParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPresentationForInteractionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
