// Code generated by go-swagger; DO NOT EDIT.

package policy_service

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

// NewGetPolicyLogsParams creates a new GetPolicyLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPolicyLogsParams() *GetPolicyLogsParams {
	return &GetPolicyLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyLogsParamsWithTimeout creates a new GetPolicyLogsParams object
// with the ability to set a timeout on a request.
func NewGetPolicyLogsParamsWithTimeout(timeout time.Duration) *GetPolicyLogsParams {
	return &GetPolicyLogsParams{
		timeout: timeout,
	}
}

// NewGetPolicyLogsParamsWithContext creates a new GetPolicyLogsParams object
// with the ability to set a context for a request.
func NewGetPolicyLogsParamsWithContext(ctx context.Context) *GetPolicyLogsParams {
	return &GetPolicyLogsParams{
		Context: ctx,
	}
}

// NewGetPolicyLogsParamsWithHTTPClient creates a new GetPolicyLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPolicyLogsParamsWithHTTPClient(client *http.Client) *GetPolicyLogsParams {
	return &GetPolicyLogsParams{
		HTTPClient: client,
	}
}

/*
GetPolicyLogsParams contains all the parameters to send to the API endpoint

	for the get policy logs operation.

	Typically these are written to a http.Request.
*/
type GetPolicyLogsParams struct {

	// Conditions.
	Conditions []string

	// ID.
	ID string

	// MaxResult.
	//
	// Format: int32
	// Default: 2147483647
	MaxResult int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get policy logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyLogsParams) WithDefaults() *GetPolicyLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get policy logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyLogsParams) SetDefaults() {
	var (
		maxResultDefault = int32(2.147483647e+09)
	)

	val := GetPolicyLogsParams{
		MaxResult: maxResultDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get policy logs params
func (o *GetPolicyLogsParams) WithTimeout(timeout time.Duration) *GetPolicyLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy logs params
func (o *GetPolicyLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy logs params
func (o *GetPolicyLogsParams) WithContext(ctx context.Context) *GetPolicyLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy logs params
func (o *GetPolicyLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy logs params
func (o *GetPolicyLogsParams) WithHTTPClient(client *http.Client) *GetPolicyLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy logs params
func (o *GetPolicyLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConditions adds the conditions to the get policy logs params
func (o *GetPolicyLogsParams) WithConditions(conditions []string) *GetPolicyLogsParams {
	o.SetConditions(conditions)
	return o
}

// SetConditions adds the conditions to the get policy logs params
func (o *GetPolicyLogsParams) SetConditions(conditions []string) {
	o.Conditions = conditions
}

// WithID adds the id to the get policy logs params
func (o *GetPolicyLogsParams) WithID(id string) *GetPolicyLogsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get policy logs params
func (o *GetPolicyLogsParams) SetID(id string) {
	o.ID = id
}

// WithMaxResult adds the maxResult to the get policy logs params
func (o *GetPolicyLogsParams) WithMaxResult(maxResult int32) *GetPolicyLogsParams {
	o.SetMaxResult(maxResult)
	return o
}

// SetMaxResult adds the maxResult to the get policy logs params
func (o *GetPolicyLogsParams) SetMaxResult(maxResult int32) {
	o.MaxResult = maxResult
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Conditions != nil {

		// binding items for conditions
		joinedConditions := o.bindParamConditions(reg)

		// query array param conditions
		if err := r.SetQueryParam("conditions", joinedConditions...); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param maxResult
	qrMaxResult := o.MaxResult
	qMaxResult := swag.FormatInt32(qrMaxResult)
	if qMaxResult != "" {

		if err := r.SetQueryParam("maxResult", qMaxResult); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetPolicyLogs binds the parameter conditions
func (o *GetPolicyLogsParams) bindParamConditions(formats strfmt.Registry) []string {
	conditionsIR := o.Conditions

	var conditionsIC []string
	for _, conditionsIIR := range conditionsIR { // explode []string

		conditionsIIV := conditionsIIR // string as string
		conditionsIC = append(conditionsIC, conditionsIIV)
	}

	// items.CollectionFormat: "multi"
	conditionsIS := swag.JoinByFormat(conditionsIC, "multi")

	return conditionsIS
}
