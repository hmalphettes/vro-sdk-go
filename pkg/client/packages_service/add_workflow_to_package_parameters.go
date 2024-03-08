// Code generated by go-swagger; DO NOT EDIT.

package packages_service

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

// NewAddWorkflowToPackageParams creates a new AddWorkflowToPackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddWorkflowToPackageParams() *AddWorkflowToPackageParams {
	return &AddWorkflowToPackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddWorkflowToPackageParamsWithTimeout creates a new AddWorkflowToPackageParams object
// with the ability to set a timeout on a request.
func NewAddWorkflowToPackageParamsWithTimeout(timeout time.Duration) *AddWorkflowToPackageParams {
	return &AddWorkflowToPackageParams{
		timeout: timeout,
	}
}

// NewAddWorkflowToPackageParamsWithContext creates a new AddWorkflowToPackageParams object
// with the ability to set a context for a request.
func NewAddWorkflowToPackageParamsWithContext(ctx context.Context) *AddWorkflowToPackageParams {
	return &AddWorkflowToPackageParams{
		Context: ctx,
	}
}

// NewAddWorkflowToPackageParamsWithHTTPClient creates a new AddWorkflowToPackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddWorkflowToPackageParamsWithHTTPClient(client *http.Client) *AddWorkflowToPackageParams {
	return &AddWorkflowToPackageParams{
		HTTPClient: client,
	}
}

/*
AddWorkflowToPackageParams contains all the parameters to send to the API endpoint

	for the add workflow to package operation.

	Typically these are written to a http.Request.
*/
type AddWorkflowToPackageParams struct {

	// PackageName.
	PackageName string

	// WorkflowID.
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add workflow to package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWorkflowToPackageParams) WithDefaults() *AddWorkflowToPackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add workflow to package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddWorkflowToPackageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add workflow to package params
func (o *AddWorkflowToPackageParams) WithTimeout(timeout time.Duration) *AddWorkflowToPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add workflow to package params
func (o *AddWorkflowToPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add workflow to package params
func (o *AddWorkflowToPackageParams) WithContext(ctx context.Context) *AddWorkflowToPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add workflow to package params
func (o *AddWorkflowToPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add workflow to package params
func (o *AddWorkflowToPackageParams) WithHTTPClient(client *http.Client) *AddWorkflowToPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add workflow to package params
func (o *AddWorkflowToPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the add workflow to package params
func (o *AddWorkflowToPackageParams) WithPackageName(packageName string) *AddWorkflowToPackageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the add workflow to package params
func (o *AddWorkflowToPackageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithWorkflowID adds the workflowID to the add workflow to package params
func (o *AddWorkflowToPackageParams) WithWorkflowID(workflowID string) *AddWorkflowToPackageParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the add workflow to package params
func (o *AddWorkflowToPackageParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *AddWorkflowToPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
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
