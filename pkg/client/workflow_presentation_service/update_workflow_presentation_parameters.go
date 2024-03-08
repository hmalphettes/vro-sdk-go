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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateWorkflowPresentationParams creates a new UpdateWorkflowPresentationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWorkflowPresentationParams() *UpdateWorkflowPresentationParams {
	return &UpdateWorkflowPresentationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWorkflowPresentationParamsWithTimeout creates a new UpdateWorkflowPresentationParams object
// with the ability to set a timeout on a request.
func NewUpdateWorkflowPresentationParamsWithTimeout(timeout time.Duration) *UpdateWorkflowPresentationParams {
	return &UpdateWorkflowPresentationParams{
		timeout: timeout,
	}
}

// NewUpdateWorkflowPresentationParamsWithContext creates a new UpdateWorkflowPresentationParams object
// with the ability to set a context for a request.
func NewUpdateWorkflowPresentationParamsWithContext(ctx context.Context) *UpdateWorkflowPresentationParams {
	return &UpdateWorkflowPresentationParams{
		Context: ctx,
	}
}

// NewUpdateWorkflowPresentationParamsWithHTTPClient creates a new UpdateWorkflowPresentationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWorkflowPresentationParamsWithHTTPClient(client *http.Client) *UpdateWorkflowPresentationParams {
	return &UpdateWorkflowPresentationParams{
		HTTPClient: client,
	}
}

/*
UpdateWorkflowPresentationParams contains all the parameters to send to the API endpoint

	for the update workflow presentation operation.

	Typically these are written to a http.Request.
*/
type UpdateWorkflowPresentationParams struct {

	// Body.
	Body *models.ExecutionContext

	// ExecutionID.
	ExecutionID string

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update workflow presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkflowPresentationParams) WithDefaults() *UpdateWorkflowPresentationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update workflow presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkflowPresentationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) WithTimeout(timeout time.Duration) *UpdateWorkflowPresentationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) WithContext(ctx context.Context) *UpdateWorkflowPresentationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) WithHTTPClient(client *http.Client) *UpdateWorkflowPresentationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) WithBody(body *models.ExecutionContext) *UpdateWorkflowPresentationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) SetBody(body *models.ExecutionContext) {
	o.Body = body
}

// WithExecutionID adds the executionID to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) WithExecutionID(executionID string) *UpdateWorkflowPresentationParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithWorkflowID adds the workflowID to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) WithWorkflowID(workflowID string) *UpdateWorkflowPresentationParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the update workflow presentation params
func (o *UpdateWorkflowPresentationParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWorkflowPresentationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
