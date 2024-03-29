// Code generated by go-swagger; DO NOT EDIT.

package category_service

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

// NewGetCategoryPermissionRuleParams creates a new GetCategoryPermissionRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCategoryPermissionRuleParams() *GetCategoryPermissionRuleParams {
	return &GetCategoryPermissionRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCategoryPermissionRuleParamsWithTimeout creates a new GetCategoryPermissionRuleParams object
// with the ability to set a timeout on a request.
func NewGetCategoryPermissionRuleParamsWithTimeout(timeout time.Duration) *GetCategoryPermissionRuleParams {
	return &GetCategoryPermissionRuleParams{
		timeout: timeout,
	}
}

// NewGetCategoryPermissionRuleParamsWithContext creates a new GetCategoryPermissionRuleParams object
// with the ability to set a context for a request.
func NewGetCategoryPermissionRuleParamsWithContext(ctx context.Context) *GetCategoryPermissionRuleParams {
	return &GetCategoryPermissionRuleParams{
		Context: ctx,
	}
}

// NewGetCategoryPermissionRuleParamsWithHTTPClient creates a new GetCategoryPermissionRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCategoryPermissionRuleParamsWithHTTPClient(client *http.Client) *GetCategoryPermissionRuleParams {
	return &GetCategoryPermissionRuleParams{
		HTTPClient: client,
	}
}

/*
GetCategoryPermissionRuleParams contains all the parameters to send to the API endpoint

	for the get category permission rule operation.

	Typically these are written to a http.Request.
*/
type GetCategoryPermissionRuleParams struct {

	// ID.
	ID string

	// RuleID.
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get category permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCategoryPermissionRuleParams) WithDefaults() *GetCategoryPermissionRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get category permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCategoryPermissionRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) WithTimeout(timeout time.Duration) *GetCategoryPermissionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) WithContext(ctx context.Context) *GetCategoryPermissionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) WithHTTPClient(client *http.Client) *GetCategoryPermissionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) WithID(id string) *GetCategoryPermissionRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) SetID(id string) {
	o.ID = id
}

// WithRuleID adds the ruleID to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) WithRuleID(ruleID string) *GetCategoryPermissionRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the get category permission rule params
func (o *GetCategoryPermissionRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCategoryPermissionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
