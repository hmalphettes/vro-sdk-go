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

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// NewChangeSlotEntryParams creates a new ChangeSlotEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeSlotEntryParams() *ChangeSlotEntryParams {
	return &ChangeSlotEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeSlotEntryParamsWithTimeout creates a new ChangeSlotEntryParams object
// with the ability to set a timeout on a request.
func NewChangeSlotEntryParamsWithTimeout(timeout time.Duration) *ChangeSlotEntryParams {
	return &ChangeSlotEntryParams{
		timeout: timeout,
	}
}

// NewChangeSlotEntryParamsWithContext creates a new ChangeSlotEntryParams object
// with the ability to set a context for a request.
func NewChangeSlotEntryParamsWithContext(ctx context.Context) *ChangeSlotEntryParams {
	return &ChangeSlotEntryParams{
		Context: ctx,
	}
}

// NewChangeSlotEntryParamsWithHTTPClient creates a new ChangeSlotEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeSlotEntryParamsWithHTTPClient(client *http.Client) *ChangeSlotEntryParams {
	return &ChangeSlotEntryParams{
		HTTPClient: client,
	}
}

/*
ChangeSlotEntryParams contains all the parameters to send to the API endpoint

	for the change slot entry operation.

	Typically these are written to a http.Request.
*/
type ChangeSlotEntryParams struct {

	// Body.
	Body *models.Parameter

	// ExecutionID.
	ExecutionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change slot entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeSlotEntryParams) WithDefaults() *ChangeSlotEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change slot entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeSlotEntryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change slot entry params
func (o *ChangeSlotEntryParams) WithTimeout(timeout time.Duration) *ChangeSlotEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change slot entry params
func (o *ChangeSlotEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change slot entry params
func (o *ChangeSlotEntryParams) WithContext(ctx context.Context) *ChangeSlotEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change slot entry params
func (o *ChangeSlotEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change slot entry params
func (o *ChangeSlotEntryParams) WithHTTPClient(client *http.Client) *ChangeSlotEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change slot entry params
func (o *ChangeSlotEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change slot entry params
func (o *ChangeSlotEntryParams) WithBody(body *models.Parameter) *ChangeSlotEntryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change slot entry params
func (o *ChangeSlotEntryParams) SetBody(body *models.Parameter) {
	o.Body = body
}

// WithExecutionID adds the executionID to the change slot entry params
func (o *ChangeSlotEntryParams) WithExecutionID(executionID string) *ChangeSlotEntryParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the change slot entry params
func (o *ChangeSlotEntryParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeSlotEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
