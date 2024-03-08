// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// NewInstallPluginDynamicallyParams creates a new InstallPluginDynamicallyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstallPluginDynamicallyParams() *InstallPluginDynamicallyParams {
	return &InstallPluginDynamicallyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstallPluginDynamicallyParamsWithTimeout creates a new InstallPluginDynamicallyParams object
// with the ability to set a timeout on a request.
func NewInstallPluginDynamicallyParamsWithTimeout(timeout time.Duration) *InstallPluginDynamicallyParams {
	return &InstallPluginDynamicallyParams{
		timeout: timeout,
	}
}

// NewInstallPluginDynamicallyParamsWithContext creates a new InstallPluginDynamicallyParams object
// with the ability to set a context for a request.
func NewInstallPluginDynamicallyParamsWithContext(ctx context.Context) *InstallPluginDynamicallyParams {
	return &InstallPluginDynamicallyParams{
		Context: ctx,
	}
}

// NewInstallPluginDynamicallyParamsWithHTTPClient creates a new InstallPluginDynamicallyParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstallPluginDynamicallyParamsWithHTTPClient(client *http.Client) *InstallPluginDynamicallyParams {
	return &InstallPluginDynamicallyParams{
		HTTPClient: client,
	}
}

/*
InstallPluginDynamicallyParams contains all the parameters to send to the API endpoint

	for the install plugin dynamically operation.

	Typically these are written to a http.Request.
*/
type InstallPluginDynamicallyParams struct {

	// File.
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the install plugin dynamically params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallPluginDynamicallyParams) WithDefaults() *InstallPluginDynamicallyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the install plugin dynamically params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallPluginDynamicallyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) WithTimeout(timeout time.Duration) *InstallPluginDynamicallyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) WithContext(ctx context.Context) *InstallPluginDynamicallyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) WithHTTPClient(client *http.Client) *InstallPluginDynamicallyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) WithFile(file runtime.NamedReadCloser) *InstallPluginDynamicallyParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the install plugin dynamically params
func (o *InstallPluginDynamicallyParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *InstallPluginDynamicallyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
