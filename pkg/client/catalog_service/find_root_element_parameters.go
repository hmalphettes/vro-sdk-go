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

// NewFindRootElementParams creates a new FindRootElementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindRootElementParams() *FindRootElementParams {
	return &FindRootElementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindRootElementParamsWithTimeout creates a new FindRootElementParams object
// with the ability to set a timeout on a request.
func NewFindRootElementParamsWithTimeout(timeout time.Duration) *FindRootElementParams {
	return &FindRootElementParams{
		timeout: timeout,
	}
}

// NewFindRootElementParamsWithContext creates a new FindRootElementParams object
// with the ability to set a context for a request.
func NewFindRootElementParamsWithContext(ctx context.Context) *FindRootElementParams {
	return &FindRootElementParams{
		Context: ctx,
	}
}

// NewFindRootElementParamsWithHTTPClient creates a new FindRootElementParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindRootElementParamsWithHTTPClient(client *http.Client) *FindRootElementParams {
	return &FindRootElementParams{
		HTTPClient: client,
	}
}

/*
FindRootElementParams contains all the parameters to send to the API endpoint

	for the find root element operation.

	Typically these are written to a http.Request.
*/
type FindRootElementParams struct {

	// Keys.
	Keys []string

	// Namespace.
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find root element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindRootElementParams) WithDefaults() *FindRootElementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find root element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindRootElementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find root element params
func (o *FindRootElementParams) WithTimeout(timeout time.Duration) *FindRootElementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find root element params
func (o *FindRootElementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find root element params
func (o *FindRootElementParams) WithContext(ctx context.Context) *FindRootElementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find root element params
func (o *FindRootElementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find root element params
func (o *FindRootElementParams) WithHTTPClient(client *http.Client) *FindRootElementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find root element params
func (o *FindRootElementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeys adds the keys to the find root element params
func (o *FindRootElementParams) WithKeys(keys []string) *FindRootElementParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the find root element params
func (o *FindRootElementParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithNamespace adds the namespace to the find root element params
func (o *FindRootElementParams) WithNamespace(namespace string) *FindRootElementParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the find root element params
func (o *FindRootElementParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *FindRootElementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Keys != nil {

		// binding items for keys
		joinedKeys := o.bindParamKeys(reg)

		// query array param keys
		if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFindRootElement binds the parameter keys
func (o *FindRootElementParams) bindParamKeys(formats strfmt.Registry) []string {
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
