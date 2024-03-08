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

// NewGetAllExecutionsForParams creates a new GetAllExecutionsForParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllExecutionsForParams() *GetAllExecutionsForParams {
	return &GetAllExecutionsForParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllExecutionsForParamsWithTimeout creates a new GetAllExecutionsForParams object
// with the ability to set a timeout on a request.
func NewGetAllExecutionsForParamsWithTimeout(timeout time.Duration) *GetAllExecutionsForParams {
	return &GetAllExecutionsForParams{
		timeout: timeout,
	}
}

// NewGetAllExecutionsForParamsWithContext creates a new GetAllExecutionsForParams object
// with the ability to set a context for a request.
func NewGetAllExecutionsForParamsWithContext(ctx context.Context) *GetAllExecutionsForParams {
	return &GetAllExecutionsForParams{
		Context: ctx,
	}
}

// NewGetAllExecutionsForParamsWithHTTPClient creates a new GetAllExecutionsForParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllExecutionsForParamsWithHTTPClient(client *http.Client) *GetAllExecutionsForParams {
	return &GetAllExecutionsForParams{
		HTTPClient: client,
	}
}

/*
GetAllExecutionsForParams contains all the parameters to send to the API endpoint

	for the get all executions for operation.

	Typically these are written to a http.Request.
*/
type GetAllExecutionsForParams struct {

	// Conditions.
	Conditions []string

	// Keys.
	Keys []string

	// MaxResult.
	//
	// Format: int32
	// Default: 2147483647
	MaxResult int32

	// SortOrders.
	SortOrders []string

	// StartIndex.
	//
	// Format: int32
	StartIndex int32

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all executions for params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllExecutionsForParams) WithDefaults() *GetAllExecutionsForParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all executions for params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllExecutionsForParams) SetDefaults() {
	var (
		maxResultDefault = int32(2.147483647e+09)

		startIndexDefault = int32(0)
	)

	val := GetAllExecutionsForParams{
		MaxResult:  maxResultDefault,
		StartIndex: startIndexDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all executions for params
func (o *GetAllExecutionsForParams) WithTimeout(timeout time.Duration) *GetAllExecutionsForParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all executions for params
func (o *GetAllExecutionsForParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all executions for params
func (o *GetAllExecutionsForParams) WithContext(ctx context.Context) *GetAllExecutionsForParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all executions for params
func (o *GetAllExecutionsForParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all executions for params
func (o *GetAllExecutionsForParams) WithHTTPClient(client *http.Client) *GetAllExecutionsForParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all executions for params
func (o *GetAllExecutionsForParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConditions adds the conditions to the get all executions for params
func (o *GetAllExecutionsForParams) WithConditions(conditions []string) *GetAllExecutionsForParams {
	o.SetConditions(conditions)
	return o
}

// SetConditions adds the conditions to the get all executions for params
func (o *GetAllExecutionsForParams) SetConditions(conditions []string) {
	o.Conditions = conditions
}

// WithKeys adds the keys to the get all executions for params
func (o *GetAllExecutionsForParams) WithKeys(keys []string) *GetAllExecutionsForParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the get all executions for params
func (o *GetAllExecutionsForParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithMaxResult adds the maxResult to the get all executions for params
func (o *GetAllExecutionsForParams) WithMaxResult(maxResult int32) *GetAllExecutionsForParams {
	o.SetMaxResult(maxResult)
	return o
}

// SetMaxResult adds the maxResult to the get all executions for params
func (o *GetAllExecutionsForParams) SetMaxResult(maxResult int32) {
	o.MaxResult = maxResult
}

// WithSortOrders adds the sortOrders to the get all executions for params
func (o *GetAllExecutionsForParams) WithSortOrders(sortOrders []string) *GetAllExecutionsForParams {
	o.SetSortOrders(sortOrders)
	return o
}

// SetSortOrders adds the sortOrders to the get all executions for params
func (o *GetAllExecutionsForParams) SetSortOrders(sortOrders []string) {
	o.SortOrders = sortOrders
}

// WithStartIndex adds the startIndex to the get all executions for params
func (o *GetAllExecutionsForParams) WithStartIndex(startIndex int32) *GetAllExecutionsForParams {
	o.SetStartIndex(startIndex)
	return o
}

// SetStartIndex adds the startIndex to the get all executions for params
func (o *GetAllExecutionsForParams) SetStartIndex(startIndex int32) {
	o.StartIndex = startIndex
}

// WithWorkflowID adds the workflowID to the get all executions for params
func (o *GetAllExecutionsForParams) WithWorkflowID(workflowID string) *GetAllExecutionsForParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the get all executions for params
func (o *GetAllExecutionsForParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllExecutionsForParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Keys != nil {

		// binding items for keys
		joinedKeys := o.bindParamKeys(reg)

		// query array param keys
		if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
			return err
		}
	}

	// query param maxResult
	qrMaxResult := o.MaxResult
	qMaxResult := swag.FormatInt32(qrMaxResult)
	if qMaxResult != "" {

		if err := r.SetQueryParam("maxResult", qMaxResult); err != nil {
			return err
		}
	}

	if o.SortOrders != nil {

		// binding items for sortOrders
		joinedSortOrders := o.bindParamSortOrders(reg)

		// query array param sortOrders
		if err := r.SetQueryParam("sortOrders", joinedSortOrders...); err != nil {
			return err
		}
	}

	// query param startIndex
	qrStartIndex := o.StartIndex
	qStartIndex := swag.FormatInt32(qrStartIndex)
	if qStartIndex != "" {

		if err := r.SetQueryParam("startIndex", qStartIndex); err != nil {
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

// bindParamGetAllExecutionsFor binds the parameter conditions
func (o *GetAllExecutionsForParams) bindParamConditions(formats strfmt.Registry) []string {
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

// bindParamGetAllExecutionsFor binds the parameter keys
func (o *GetAllExecutionsForParams) bindParamKeys(formats strfmt.Registry) []string {
	keysIR := o.Keys

	var keysIC []string
	for _, keysIIR := range keysIR { // explode []string

		keysIIV := keysIIR // string as string
		keysIC = append(keysIC, keysIIV)
	}

	// items.CollectionFormat: "multi"
	keysIS := swag.JoinByFormat(keysIC, "multi")

	return keysIS
}

// bindParamGetAllExecutionsFor binds the parameter sortOrders
func (o *GetAllExecutionsForParams) bindParamSortOrders(formats strfmt.Registry) []string {
	sortOrdersIR := o.SortOrders

	var sortOrdersIC []string
	for _, sortOrdersIIR := range sortOrdersIR { // explode []string

		sortOrdersIIV := sortOrdersIIR // string as string
		sortOrdersIC = append(sortOrdersIC, sortOrdersIIV)
	}

	// items.CollectionFormat: "multi"
	sortOrdersIS := swag.JoinByFormat(sortOrdersIC, "multi")

	return sortOrdersIS
}