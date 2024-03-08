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

// NewStartPresentationParams creates a new StartPresentationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartPresentationParams() *StartPresentationParams {
	return &StartPresentationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartPresentationParamsWithTimeout creates a new StartPresentationParams object
// with the ability to set a timeout on a request.
func NewStartPresentationParamsWithTimeout(timeout time.Duration) *StartPresentationParams {
	return &StartPresentationParams{
		timeout: timeout,
	}
}

// NewStartPresentationParamsWithContext creates a new StartPresentationParams object
// with the ability to set a context for a request.
func NewStartPresentationParamsWithContext(ctx context.Context) *StartPresentationParams {
	return &StartPresentationParams{
		Context: ctx,
	}
}

// NewStartPresentationParamsWithHTTPClient creates a new StartPresentationParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartPresentationParamsWithHTTPClient(client *http.Client) *StartPresentationParams {
	return &StartPresentationParams{
		HTTPClient: client,
	}
}

/*
StartPresentationParams contains all the parameters to send to the API endpoint

	for the start presentation operation.

	Typically these are written to a http.Request.
*/
type StartPresentationParams struct {

	// Body.
	Body *models.ExecutionContext

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartPresentationParams) WithDefaults() *StartPresentationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start presentation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartPresentationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start presentation params
func (o *StartPresentationParams) WithTimeout(timeout time.Duration) *StartPresentationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start presentation params
func (o *StartPresentationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start presentation params
func (o *StartPresentationParams) WithContext(ctx context.Context) *StartPresentationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start presentation params
func (o *StartPresentationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start presentation params
func (o *StartPresentationParams) WithHTTPClient(client *http.Client) *StartPresentationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start presentation params
func (o *StartPresentationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start presentation params
func (o *StartPresentationParams) WithBody(body *models.ExecutionContext) *StartPresentationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start presentation params
func (o *StartPresentationParams) SetBody(body *models.ExecutionContext) {
	o.Body = body
}

// WithWorkflowID adds the workflowID to the start presentation params
func (o *StartPresentationParams) WithWorkflowID(workflowID string) *StartPresentationParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the start presentation params
func (o *StartPresentationParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *StartPresentationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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