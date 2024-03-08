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

// NewDownloadIconForModuleParams creates a new DownloadIconForModuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadIconForModuleParams() *DownloadIconForModuleParams {
	return &DownloadIconForModuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadIconForModuleParamsWithTimeout creates a new DownloadIconForModuleParams object
// with the ability to set a timeout on a request.
func NewDownloadIconForModuleParamsWithTimeout(timeout time.Duration) *DownloadIconForModuleParams {
	return &DownloadIconForModuleParams{
		timeout: timeout,
	}
}

// NewDownloadIconForModuleParamsWithContext creates a new DownloadIconForModuleParams object
// with the ability to set a context for a request.
func NewDownloadIconForModuleParamsWithContext(ctx context.Context) *DownloadIconForModuleParams {
	return &DownloadIconForModuleParams{
		Context: ctx,
	}
}

// NewDownloadIconForModuleParamsWithHTTPClient creates a new DownloadIconForModuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadIconForModuleParamsWithHTTPClient(client *http.Client) *DownloadIconForModuleParams {
	return &DownloadIconForModuleParams{
		HTTPClient: client,
	}
}

/*
DownloadIconForModuleParams contains all the parameters to send to the API endpoint

	for the download icon for module operation.

	Typically these are written to a http.Request.
*/
type DownloadIconForModuleParams struct {

	// Namespace.
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download icon for module params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadIconForModuleParams) WithDefaults() *DownloadIconForModuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download icon for module params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadIconForModuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download icon for module params
func (o *DownloadIconForModuleParams) WithTimeout(timeout time.Duration) *DownloadIconForModuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download icon for module params
func (o *DownloadIconForModuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download icon for module params
func (o *DownloadIconForModuleParams) WithContext(ctx context.Context) *DownloadIconForModuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download icon for module params
func (o *DownloadIconForModuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download icon for module params
func (o *DownloadIconForModuleParams) WithHTTPClient(client *http.Client) *DownloadIconForModuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download icon for module params
func (o *DownloadIconForModuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the download icon for module params
func (o *DownloadIconForModuleParams) WithNamespace(namespace string) *DownloadIconForModuleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the download icon for module params
func (o *DownloadIconForModuleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadIconForModuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
