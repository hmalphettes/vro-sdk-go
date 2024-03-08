// Code generated by go-swagger; DO NOT EDIT.

package configuration_service

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

// NewUpdateConfigurationPermissionRuleParams creates a new UpdateConfigurationPermissionRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConfigurationPermissionRuleParams() *UpdateConfigurationPermissionRuleParams {
	return &UpdateConfigurationPermissionRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigurationPermissionRuleParamsWithTimeout creates a new UpdateConfigurationPermissionRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateConfigurationPermissionRuleParamsWithTimeout(timeout time.Duration) *UpdateConfigurationPermissionRuleParams {
	return &UpdateConfigurationPermissionRuleParams{
		timeout: timeout,
	}
}

// NewUpdateConfigurationPermissionRuleParamsWithContext creates a new UpdateConfigurationPermissionRuleParams object
// with the ability to set a context for a request.
func NewUpdateConfigurationPermissionRuleParamsWithContext(ctx context.Context) *UpdateConfigurationPermissionRuleParams {
	return &UpdateConfigurationPermissionRuleParams{
		Context: ctx,
	}
}

// NewUpdateConfigurationPermissionRuleParamsWithHTTPClient creates a new UpdateConfigurationPermissionRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConfigurationPermissionRuleParamsWithHTTPClient(client *http.Client) *UpdateConfigurationPermissionRuleParams {
	return &UpdateConfigurationPermissionRuleParams{
		HTTPClient: client,
	}
}

/*
UpdateConfigurationPermissionRuleParams contains all the parameters to send to the API endpoint

	for the update configuration permission rule operation.

	Typically these are written to a http.Request.
*/
type UpdateConfigurationPermissionRuleParams struct {

	// Body.
	Body *models.Permission

	// ID.
	ID string

	// RuleID.
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update configuration permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfigurationPermissionRuleParams) WithDefaults() *UpdateConfigurationPermissionRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update configuration permission rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConfigurationPermissionRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) WithTimeout(timeout time.Duration) *UpdateConfigurationPermissionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) WithContext(ctx context.Context) *UpdateConfigurationPermissionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) WithHTTPClient(client *http.Client) *UpdateConfigurationPermissionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) WithBody(body *models.Permission) *UpdateConfigurationPermissionRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) SetBody(body *models.Permission) {
	o.Body = body
}

// WithID adds the id to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) WithID(id string) *UpdateConfigurationPermissionRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) SetID(id string) {
	o.ID = id
}

// WithRuleID adds the ruleID to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) WithRuleID(ruleID string) *UpdateConfigurationPermissionRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the update configuration permission rule params
func (o *UpdateConfigurationPermissionRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigurationPermissionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
