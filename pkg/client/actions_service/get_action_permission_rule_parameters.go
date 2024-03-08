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
)

// NewGetActionPermissionRuleParams creates a new GetActionPermissionRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetActionPermissionRuleParams() *GetActionPermissionRuleParams {
	return &GetActionPermissionRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionPermissionRuleParamsWithTimeout creates a new GetActionPermissionRuleParams object
// with the ability to set a timeout on a request.
func NewGetActionPermissionRuleParamsWithTimeout(timeout time.Duration) *GetActionPermissionRuleParams {
	return &GetActionPermissionRuleParams{
		timeout: timeout,
	}
}

// NewGetActionPermissionRuleParamsWithContext creates a new GetActionPermissionRuleParams object
// with the ability to set a context for a request.
func NewGetActionPermissionRuleParamsWithContext(ctx context.Context) *GetActionPermissionRuleParams {
	return &GetActionPermissionRuleParams{
		Context: ctx,
	}
}

// NewGetActionPermissionRuleParamsWithHTTPClient creates a new GetActionPermissionRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetActionPermissionRuleParamsWithHTTPClient(client *http.Client) *GetActionPermissionRuleParams {
	return &GetActionPermissionRuleParams{
		HTTPClient: client,
	}
}

/*
GetActionPermissionRuleParams contains all the parameters to send to the API endpoint

	for the get action permission rule operation.

	Typically these are written to a http.Request.
*/
type GetActionPermissionRuleParams struct {

	// ID.
	ID string

	// RuleID.
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get action permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionPermissionRuleParams) WithDefaults() *GetActionPermissionRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get action permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionPermissionRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get action permission rule params
func (o *GetActionPermissionRuleParams) WithTimeout(timeout time.Duration) *GetActionPermissionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action permission rule params
func (o *GetActionPermissionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action permission rule params
func (o *GetActionPermissionRuleParams) WithContext(ctx context.Context) *GetActionPermissionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action permission rule params
func (o *GetActionPermissionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action permission rule params
func (o *GetActionPermissionRuleParams) WithHTTPClient(client *http.Client) *GetActionPermissionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action permission rule params
func (o *GetActionPermissionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get action permission rule params
func (o *GetActionPermissionRuleParams) WithID(id string) *GetActionPermissionRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get action permission rule params
func (o *GetActionPermissionRuleParams) SetID(id string) {
	o.ID = id
}

// WithRuleID adds the ruleID to the get action permission rule params
func (o *GetActionPermissionRuleParams) WithRuleID(ruleID string) *GetActionPermissionRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the get action permission rule params
func (o *GetActionPermissionRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionPermissionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
