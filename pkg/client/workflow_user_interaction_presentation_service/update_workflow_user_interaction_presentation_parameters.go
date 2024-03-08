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

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// NewUpdateWorkflowUserInteractionPresentationParams creates a new UpdateWorkflowUserInteractionPresentationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWorkflowUserInteractionPresentationParams() *UpdateWorkflowUserInteractionPresentationParams {
	return &UpdateWorkflowUserInteractionPresentationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWorkflowUserInteractionPresentationParamsWithTimeout creates a new UpdateWorkflowUserInteractionPresentationParams object
// with the ability to set a timeout on a request.
func NewUpdateWorkflowUserInteractionPresentationParamsWithTimeout(timeout time.Duration) *UpdateWorkflowUserInteractionPresentationParams {
	return &UpdateWorkflowUserInteractionPresentationParams{
		timeout: timeout,
	}
}

// NewUpdateWorkflowUserInteractionPresentationParamsWithContext creates a new UpdateWorkflowUserInteractionPresentationParams object
// with the ability to set a context for a request.
func NewUpdateWorkflowUserInteractionPresentationParamsWithContext(ctx context.Context) *UpdateWorkflowUserInteractionPresentationParams {
	return &UpdateWorkflowUserInteractionPresentationParams{
		Context: ctx,
	}
}

// NewUpdateWorkflowUserInteractionPresentationParamsWithHTTPClient creates a new UpdateWorkflowUserInteractionPresentationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWorkflowUserInteractionPresentationParamsWithHTTPClient(client *http.Client) *UpdateWorkflowUserInteractionPresentationParams {
	return &UpdateWorkflowUserInteractionPresentationParams{
		HTTPClient: client,
	}
}

/*
UpdateWorkflowUserInteractionPresentationParams contains all the parameters to send to the API endpoint

	for the update workflow user interaction presentation operation.

	Typically these are written to a http.Request.
*/
type UpdateWorkflowUserInteractionPresentationParams struct {

	// Body.
	Body *models.ExecutionContext

	// ExecutionID.
	ExecutionID string

	// PresentationExecutionID.
	PresentationExecutionID string

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update workflow user interaction presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkflowUserInteractionPresentationParams) WithDefaults() *UpdateWorkflowUserInteractionPresentationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update workflow user interaction presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkflowUserInteractionPresentationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) WithTimeout(timeout time.Duration) *UpdateWorkflowUserInteractionPresentationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) WithContext(ctx context.Context) *UpdateWorkflowUserInteractionPresentationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) WithHTTPClient(client *http.Client) *UpdateWorkflowUserInteractionPresentationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) WithBody(body *models.ExecutionContext) *UpdateWorkflowUserInteractionPresentationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) SetBody(body *models.ExecutionContext) {
	o.Body = body
}

// WithExecutionID adds the executionID to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) WithExecutionID(executionID string) *UpdateWorkflowUserInteractionPresentationParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithPresentationExecutionID adds the presentationExecutionID to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) WithPresentationExecutionID(presentationExecutionID string) *UpdateWorkflowUserInteractionPresentationParams {
	o.SetPresentationExecutionID(presentationExecutionID)
	return o
}

// SetPresentationExecutionID adds the presentationExecutionId to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) SetPresentationExecutionID(presentationExecutionID string) {
	o.PresentationExecutionID = presentationExecutionID
}

// WithWorkflowID adds the workflowID to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) WithWorkflowID(workflowID string) *UpdateWorkflowUserInteractionPresentationParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the update workflow user interaction presentation params
func (o *UpdateWorkflowUserInteractionPresentationParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWorkflowUserInteractionPresentationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param presentationExecutionId
	if err := r.SetPathParam("presentationExecutionId", o.PresentationExecutionID); err != nil {
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
