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
	"github.com/go-openapi/swag"
)

// NewSetWorkflowBreakpointParams creates a new SetWorkflowBreakpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetWorkflowBreakpointParams() *SetWorkflowBreakpointParams {
	return &SetWorkflowBreakpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetWorkflowBreakpointParamsWithTimeout creates a new SetWorkflowBreakpointParams object
// with the ability to set a timeout on a request.
func NewSetWorkflowBreakpointParamsWithTimeout(timeout time.Duration) *SetWorkflowBreakpointParams {
	return &SetWorkflowBreakpointParams{
		timeout: timeout,
	}
}

// NewSetWorkflowBreakpointParamsWithContext creates a new SetWorkflowBreakpointParams object
// with the ability to set a context for a request.
func NewSetWorkflowBreakpointParamsWithContext(ctx context.Context) *SetWorkflowBreakpointParams {
	return &SetWorkflowBreakpointParams{
		Context: ctx,
	}
}

// NewSetWorkflowBreakpointParamsWithHTTPClient creates a new SetWorkflowBreakpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetWorkflowBreakpointParamsWithHTTPClient(client *http.Client) *SetWorkflowBreakpointParams {
	return &SetWorkflowBreakpointParams{
		HTTPClient: client,
	}
}

/*
SetWorkflowBreakpointParams contains all the parameters to send to the API endpoint

	for the set workflow breakpoint operation.

	Typically these are written to a http.Request.
*/
type SetWorkflowBreakpointParams struct {

	// ElementName.
	ElementName string

	// LineNumber.
	//
	// Format: int32
	LineNumber int32

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set workflow breakpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetWorkflowBreakpointParams) WithDefaults() *SetWorkflowBreakpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set workflow breakpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetWorkflowBreakpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) WithTimeout(timeout time.Duration) *SetWorkflowBreakpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) WithContext(ctx context.Context) *SetWorkflowBreakpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) WithHTTPClient(client *http.Client) *SetWorkflowBreakpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithElementName adds the elementName to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) WithElementName(elementName string) *SetWorkflowBreakpointParams {
	o.SetElementName(elementName)
	return o
}

// SetElementName adds the elementName to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) SetElementName(elementName string) {
	o.ElementName = elementName
}

// WithLineNumber adds the lineNumber to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) WithLineNumber(lineNumber int32) *SetWorkflowBreakpointParams {
	o.SetLineNumber(lineNumber)
	return o
}

// SetLineNumber adds the lineNumber to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) SetLineNumber(lineNumber int32) {
	o.LineNumber = lineNumber
}

// WithWorkflowID adds the workflowID to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) WithWorkflowID(workflowID string) *SetWorkflowBreakpointParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the set workflow breakpoint params
func (o *SetWorkflowBreakpointParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *SetWorkflowBreakpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param elementName
	qrElementName := o.ElementName
	qElementName := qrElementName
	if qElementName != "" {

		if err := r.SetQueryParam("elementName", qElementName); err != nil {
			return err
		}
	}

	// query param lineNumber
	qrLineNumber := o.LineNumber
	qLineNumber := swag.FormatInt32(qrLineNumber)
	if qLineNumber != "" {

		if err := r.SetQueryParam("lineNumber", qLineNumber); err != nil {
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
