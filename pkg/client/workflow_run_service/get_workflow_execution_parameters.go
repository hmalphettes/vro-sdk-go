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
	"github.com/go-openapi/swag"
)

// NewGetWorkflowExecutionParams creates a new GetWorkflowExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowExecutionParams() *GetWorkflowExecutionParams {
	return &GetWorkflowExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowExecutionParamsWithTimeout creates a new GetWorkflowExecutionParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowExecutionParamsWithTimeout(timeout time.Duration) *GetWorkflowExecutionParams {
	return &GetWorkflowExecutionParams{
		timeout: timeout,
	}
}

// NewGetWorkflowExecutionParamsWithContext creates a new GetWorkflowExecutionParams object
// with the ability to set a context for a request.
func NewGetWorkflowExecutionParamsWithContext(ctx context.Context) *GetWorkflowExecutionParams {
	return &GetWorkflowExecutionParams{
		Context: ctx,
	}
}

// NewGetWorkflowExecutionParamsWithHTTPClient creates a new GetWorkflowExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowExecutionParamsWithHTTPClient(client *http.Client) *GetWorkflowExecutionParams {
	return &GetWorkflowExecutionParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowExecutionParams contains all the parameters to send to the API endpoint

	for the get workflow execution operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowExecutionParams struct {

	// ExecutionID.
	ExecutionID string

	// Expand.
	Expand []string

	// ShowDetails.
	ShowDetails *bool

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowExecutionParams) WithDefaults() *GetWorkflowExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowExecutionParams) SetDefaults() {
	var (
		showDetailsDefault = bool(false)
	)

	val := GetWorkflowExecutionParams{
		ShowDetails: &showDetailsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get workflow execution params
func (o *GetWorkflowExecutionParams) WithTimeout(timeout time.Duration) *GetWorkflowExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow execution params
func (o *GetWorkflowExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow execution params
func (o *GetWorkflowExecutionParams) WithContext(ctx context.Context) *GetWorkflowExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow execution params
func (o *GetWorkflowExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow execution params
func (o *GetWorkflowExecutionParams) WithHTTPClient(client *http.Client) *GetWorkflowExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow execution params
func (o *GetWorkflowExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get workflow execution params
func (o *GetWorkflowExecutionParams) WithExecutionID(executionID string) *GetWorkflowExecutionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get workflow execution params
func (o *GetWorkflowExecutionParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithExpand adds the expand to the get workflow execution params
func (o *GetWorkflowExecutionParams) WithExpand(expand []string) *GetWorkflowExecutionParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get workflow execution params
func (o *GetWorkflowExecutionParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithShowDetails adds the showDetails to the get workflow execution params
func (o *GetWorkflowExecutionParams) WithShowDetails(showDetails *bool) *GetWorkflowExecutionParams {
	o.SetShowDetails(showDetails)
	return o
}

// SetShowDetails adds the showDetails to the get workflow execution params
func (o *GetWorkflowExecutionParams) SetShowDetails(showDetails *bool) {
	o.ShowDetails = showDetails
}

// WithWorkflowID adds the workflowID to the get workflow execution params
func (o *GetWorkflowExecutionParams) WithWorkflowID(workflowID string) *GetWorkflowExecutionParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the get workflow execution params
func (o *GetWorkflowExecutionParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionId
	if err := r.SetPathParam("executionId", o.ExecutionID); err != nil {
		return err
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if o.ShowDetails != nil {

		// query param showDetails
		var qrShowDetails bool

		if o.ShowDetails != nil {
			qrShowDetails = *o.ShowDetails
		}
		qShowDetails := swag.FormatBool(qrShowDetails)
		if qShowDetails != "" {

			if err := r.SetQueryParam("showDetails", qShowDetails); err != nil {
				return err
			}
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

// bindParamGetWorkflowExecution binds the parameter expand
func (o *GetWorkflowExecutionParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
