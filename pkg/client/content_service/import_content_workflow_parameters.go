// Code generated by go-swagger; DO NOT EDIT.

package content_service

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

// NewImportContentWorkflowParams creates a new ImportContentWorkflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportContentWorkflowParams() *ImportContentWorkflowParams {
	return &ImportContentWorkflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportContentWorkflowParamsWithTimeout creates a new ImportContentWorkflowParams object
// with the ability to set a timeout on a request.
func NewImportContentWorkflowParamsWithTimeout(timeout time.Duration) *ImportContentWorkflowParams {
	return &ImportContentWorkflowParams{
		timeout: timeout,
	}
}

// NewImportContentWorkflowParamsWithContext creates a new ImportContentWorkflowParams object
// with the ability to set a context for a request.
func NewImportContentWorkflowParamsWithContext(ctx context.Context) *ImportContentWorkflowParams {
	return &ImportContentWorkflowParams{
		Context: ctx,
	}
}

// NewImportContentWorkflowParamsWithHTTPClient creates a new ImportContentWorkflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportContentWorkflowParamsWithHTTPClient(client *http.Client) *ImportContentWorkflowParams {
	return &ImportContentWorkflowParams{
		HTTPClient: client,
	}
}

/*
ImportContentWorkflowParams contains all the parameters to send to the API endpoint

	for the import content workflow operation.

	Typically these are written to a http.Request.
*/
type ImportContentWorkflowParams struct {

	// CategoryID.
	CategoryID string

	// File.
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import content workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportContentWorkflowParams) WithDefaults() *ImportContentWorkflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import content workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportContentWorkflowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import content workflow params
func (o *ImportContentWorkflowParams) WithTimeout(timeout time.Duration) *ImportContentWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import content workflow params
func (o *ImportContentWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import content workflow params
func (o *ImportContentWorkflowParams) WithContext(ctx context.Context) *ImportContentWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import content workflow params
func (o *ImportContentWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import content workflow params
func (o *ImportContentWorkflowParams) WithHTTPClient(client *http.Client) *ImportContentWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import content workflow params
func (o *ImportContentWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the import content workflow params
func (o *ImportContentWorkflowParams) WithCategoryID(categoryID string) *ImportContentWorkflowParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the import content workflow params
func (o *ImportContentWorkflowParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithFile adds the file to the import content workflow params
func (o *ImportContentWorkflowParams) WithFile(file runtime.NamedReadCloser) *ImportContentWorkflowParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the import content workflow params
func (o *ImportContentWorkflowParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *ImportContentWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param categoryId
	if err := r.SetPathParam("categoryId", o.CategoryID); err != nil {
		return err
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}