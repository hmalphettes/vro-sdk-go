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
)

// NewUntagObjectParams creates a new UntagObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUntagObjectParams() *UntagObjectParams {
	return &UntagObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUntagObjectParamsWithTimeout creates a new UntagObjectParams object
// with the ability to set a timeout on a request.
func NewUntagObjectParamsWithTimeout(timeout time.Duration) *UntagObjectParams {
	return &UntagObjectParams{
		timeout: timeout,
	}
}

// NewUntagObjectParamsWithContext creates a new UntagObjectParams object
// with the ability to set a context for a request.
func NewUntagObjectParamsWithContext(ctx context.Context) *UntagObjectParams {
	return &UntagObjectParams{
		Context: ctx,
	}
}

// NewUntagObjectParamsWithHTTPClient creates a new UntagObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewUntagObjectParamsWithHTTPClient(client *http.Client) *UntagObjectParams {
	return &UntagObjectParams{
		HTTPClient: client,
	}
}

/*
UntagObjectParams contains all the parameters to send to the API endpoint

	for the untag object operation.

	Typically these are written to a http.Request.
*/
type UntagObjectParams struct {

	// ID.
	ID string

	// Namespace.
	Namespace string

	// Tagname.
	Tagname string

	// Type.
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the untag object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UntagObjectParams) WithDefaults() *UntagObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the untag object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UntagObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the untag object params
func (o *UntagObjectParams) WithTimeout(timeout time.Duration) *UntagObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the untag object params
func (o *UntagObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the untag object params
func (o *UntagObjectParams) WithContext(ctx context.Context) *UntagObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the untag object params
func (o *UntagObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the untag object params
func (o *UntagObjectParams) WithHTTPClient(client *http.Client) *UntagObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the untag object params
func (o *UntagObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the untag object params
func (o *UntagObjectParams) WithID(id string) *UntagObjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the untag object params
func (o *UntagObjectParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the untag object params
func (o *UntagObjectParams) WithNamespace(namespace string) *UntagObjectParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the untag object params
func (o *UntagObjectParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTagname adds the tagname to the untag object params
func (o *UntagObjectParams) WithTagname(tagname string) *UntagObjectParams {
	o.SetTagname(tagname)
	return o
}

// SetTagname adds the tagname to the untag object params
func (o *UntagObjectParams) SetTagname(tagname string) {
	o.Tagname = tagname
}

// WithType adds the typeVar to the untag object params
func (o *UntagObjectParams) WithType(typeVar string) *UntagObjectParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the untag object params
func (o *UntagObjectParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *UntagObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param tagname
	if err := r.SetPathParam("tagname", o.Tagname); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
