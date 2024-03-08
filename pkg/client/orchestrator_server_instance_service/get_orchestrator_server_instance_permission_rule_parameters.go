// Code generated by go-swagger; DO NOT EDIT.

package orchestrator_server_instance_service

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

// NewGetOrchestratorServerInstancePermissionRuleParams creates a new GetOrchestratorServerInstancePermissionRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrchestratorServerInstancePermissionRuleParams() *GetOrchestratorServerInstancePermissionRuleParams {
	return &GetOrchestratorServerInstancePermissionRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrchestratorServerInstancePermissionRuleParamsWithTimeout creates a new GetOrchestratorServerInstancePermissionRuleParams object
// with the ability to set a timeout on a request.
func NewGetOrchestratorServerInstancePermissionRuleParamsWithTimeout(timeout time.Duration) *GetOrchestratorServerInstancePermissionRuleParams {
	return &GetOrchestratorServerInstancePermissionRuleParams{
		timeout: timeout,
	}
}

// NewGetOrchestratorServerInstancePermissionRuleParamsWithContext creates a new GetOrchestratorServerInstancePermissionRuleParams object
// with the ability to set a context for a request.
func NewGetOrchestratorServerInstancePermissionRuleParamsWithContext(ctx context.Context) *GetOrchestratorServerInstancePermissionRuleParams {
	return &GetOrchestratorServerInstancePermissionRuleParams{
		Context: ctx,
	}
}

// NewGetOrchestratorServerInstancePermissionRuleParamsWithHTTPClient creates a new GetOrchestratorServerInstancePermissionRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrchestratorServerInstancePermissionRuleParamsWithHTTPClient(client *http.Client) *GetOrchestratorServerInstancePermissionRuleParams {
	return &GetOrchestratorServerInstancePermissionRuleParams{
		HTTPClient: client,
	}
}

/*
GetOrchestratorServerInstancePermissionRuleParams contains all the parameters to send to the API endpoint

	for the get orchestrator server instance permission rule operation.

	Typically these are written to a http.Request.
*/
type GetOrchestratorServerInstancePermissionRuleParams struct {

	// RuleID.
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orchestrator server instance permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrchestratorServerInstancePermissionRuleParams) WithDefaults() *GetOrchestratorServerInstancePermissionRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orchestrator server instance permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrchestratorServerInstancePermissionRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) WithTimeout(timeout time.Duration) *GetOrchestratorServerInstancePermissionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) WithContext(ctx context.Context) *GetOrchestratorServerInstancePermissionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) WithHTTPClient(client *http.Client) *GetOrchestratorServerInstancePermissionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) WithRuleID(ruleID string) *GetOrchestratorServerInstancePermissionRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the get orchestrator server instance permission rule params
func (o *GetOrchestratorServerInstancePermissionRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrchestratorServerInstancePermissionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}