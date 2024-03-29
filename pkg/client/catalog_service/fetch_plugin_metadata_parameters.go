// Code generated by go-swagger; DO NOT EDIT.

package catalog_service

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

// NewFetchPluginMetadataParams creates a new FetchPluginMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFetchPluginMetadataParams() *FetchPluginMetadataParams {
	return &FetchPluginMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFetchPluginMetadataParamsWithTimeout creates a new FetchPluginMetadataParams object
// with the ability to set a timeout on a request.
func NewFetchPluginMetadataParamsWithTimeout(timeout time.Duration) *FetchPluginMetadataParams {
	return &FetchPluginMetadataParams{
		timeout: timeout,
	}
}

// NewFetchPluginMetadataParamsWithContext creates a new FetchPluginMetadataParams object
// with the ability to set a context for a request.
func NewFetchPluginMetadataParamsWithContext(ctx context.Context) *FetchPluginMetadataParams {
	return &FetchPluginMetadataParams{
		Context: ctx,
	}
}

// NewFetchPluginMetadataParamsWithHTTPClient creates a new FetchPluginMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewFetchPluginMetadataParamsWithHTTPClient(client *http.Client) *FetchPluginMetadataParams {
	return &FetchPluginMetadataParams{
		HTTPClient: client,
	}
}

/*
FetchPluginMetadataParams contains all the parameters to send to the API endpoint

	for the fetch plugin metadata operation.

	Typically these are written to a http.Request.
*/
type FetchPluginMetadataParams struct {

	// Namespace.
	Namespace string

	// ShowHiddenTypes.
	ShowHiddenTypes *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fetch plugin metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchPluginMetadataParams) WithDefaults() *FetchPluginMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fetch plugin metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FetchPluginMetadataParams) SetDefaults() {
	var (
		showHiddenTypesDefault = bool(false)
	)

	val := FetchPluginMetadataParams{
		ShowHiddenTypes: &showHiddenTypesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) WithTimeout(timeout time.Duration) *FetchPluginMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) WithContext(ctx context.Context) *FetchPluginMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) WithHTTPClient(client *http.Client) *FetchPluginMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) WithNamespace(namespace string) *FetchPluginMetadataParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithShowHiddenTypes adds the showHiddenTypes to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) WithShowHiddenTypes(showHiddenTypes *bool) *FetchPluginMetadataParams {
	o.SetShowHiddenTypes(showHiddenTypes)
	return o
}

// SetShowHiddenTypes adds the showHiddenTypes to the fetch plugin metadata params
func (o *FetchPluginMetadataParams) SetShowHiddenTypes(showHiddenTypes *bool) {
	o.ShowHiddenTypes = showHiddenTypes
}

// WriteToRequest writes these params to a swagger request
func (o *FetchPluginMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.ShowHiddenTypes != nil {

		// query param showHiddenTypes
		var qrShowHiddenTypes bool

		if o.ShowHiddenTypes != nil {
			qrShowHiddenTypes = *o.ShowHiddenTypes
		}
		qShowHiddenTypes := swag.FormatBool(qrShowHiddenTypes)
		if qShowHiddenTypes != "" {

			if err := r.SetQueryParam("showHiddenTypes", qShowHiddenTypes); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
