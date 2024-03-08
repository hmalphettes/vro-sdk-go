// Code generated by go-swagger; DO NOT EDIT.

package actions_service

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

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// NewExecuteActionByIDParams creates a new ExecuteActionByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteActionByIDParams() *ExecuteActionByIDParams {
	return &ExecuteActionByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteActionByIDParamsWithTimeout creates a new ExecuteActionByIDParams object
// with the ability to set a timeout on a request.
func NewExecuteActionByIDParamsWithTimeout(timeout time.Duration) *ExecuteActionByIDParams {
	return &ExecuteActionByIDParams{
		timeout: timeout,
	}
}

// NewExecuteActionByIDParamsWithContext creates a new ExecuteActionByIDParams object
// with the ability to set a context for a request.
func NewExecuteActionByIDParamsWithContext(ctx context.Context) *ExecuteActionByIDParams {
	return &ExecuteActionByIDParams{
		Context: ctx,
	}
}

// NewExecuteActionByIDParamsWithHTTPClient creates a new ExecuteActionByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteActionByIDParamsWithHTTPClient(client *http.Client) *ExecuteActionByIDParams {
	return &ExecuteActionByIDParams{
		HTTPClient: client,
	}
}

/*
ExecuteActionByIDParams contains all the parameters to send to the API endpoint

	for the execute action by Id operation.

	Typically these are written to a http.Request.
*/
type ExecuteActionByIDParams struct {

	// ActionID.
	ActionID string

	// Body.
	Body *models.ExecutionContext

	// Expand.
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute action by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteActionByIDParams) WithDefaults() *ExecuteActionByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute action by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteActionByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execute action by Id params
func (o *ExecuteActionByIDParams) WithTimeout(timeout time.Duration) *ExecuteActionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute action by Id params
func (o *ExecuteActionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute action by Id params
func (o *ExecuteActionByIDParams) WithContext(ctx context.Context) *ExecuteActionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute action by Id params
func (o *ExecuteActionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute action by Id params
func (o *ExecuteActionByIDParams) WithHTTPClient(client *http.Client) *ExecuteActionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute action by Id params
func (o *ExecuteActionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the execute action by Id params
func (o *ExecuteActionByIDParams) WithActionID(actionID string) *ExecuteActionByIDParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the execute action by Id params
func (o *ExecuteActionByIDParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithBody adds the body to the execute action by Id params
func (o *ExecuteActionByIDParams) WithBody(body *models.ExecutionContext) *ExecuteActionByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute action by Id params
func (o *ExecuteActionByIDParams) SetBody(body *models.ExecutionContext) {
	o.Body = body
}

// WithExpand adds the expand to the execute action by Id params
func (o *ExecuteActionByIDParams) WithExpand(expand []string) *ExecuteActionByIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the execute action by Id params
func (o *ExecuteActionByIDParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteActionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamExecuteActionByID binds the parameter expand
func (o *ExecuteActionByIDParams) bindParamExpand(formats strfmt.Registry) []string {
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
