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

// NewGetPermissionsForPackageParams creates a new GetPermissionsForPackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPermissionsForPackageParams() *GetPermissionsForPackageParams {
	return &GetPermissionsForPackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPermissionsForPackageParamsWithTimeout creates a new GetPermissionsForPackageParams object
// with the ability to set a timeout on a request.
func NewGetPermissionsForPackageParamsWithTimeout(timeout time.Duration) *GetPermissionsForPackageParams {
	return &GetPermissionsForPackageParams{
		timeout: timeout,
	}
}

// NewGetPermissionsForPackageParamsWithContext creates a new GetPermissionsForPackageParams object
// with the ability to set a context for a request.
func NewGetPermissionsForPackageParamsWithContext(ctx context.Context) *GetPermissionsForPackageParams {
	return &GetPermissionsForPackageParams{
		Context: ctx,
	}
}

// NewGetPermissionsForPackageParamsWithHTTPClient creates a new GetPermissionsForPackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPermissionsForPackageParamsWithHTTPClient(client *http.Client) *GetPermissionsForPackageParams {
	return &GetPermissionsForPackageParams{
		HTTPClient: client,
	}
}

/*
GetPermissionsForPackageParams contains all the parameters to send to the API endpoint

	for the get permissions for package operation.

	Typically these are written to a http.Request.
*/
type GetPermissionsForPackageParams struct {

	// PackageName.
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get permissions for package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionsForPackageParams) WithDefaults() *GetPermissionsForPackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get permissions for package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionsForPackageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get permissions for package params
func (o *GetPermissionsForPackageParams) WithTimeout(timeout time.Duration) *GetPermissionsForPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get permissions for package params
func (o *GetPermissionsForPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get permissions for package params
func (o *GetPermissionsForPackageParams) WithContext(ctx context.Context) *GetPermissionsForPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get permissions for package params
func (o *GetPermissionsForPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get permissions for package params
func (o *GetPermissionsForPackageParams) WithHTTPClient(client *http.Client) *GetPermissionsForPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get permissions for package params
func (o *GetPermissionsForPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the get permissions for package params
func (o *GetPermissionsForPackageParams) WithPackageName(packageName string) *GetPermissionsForPackageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the get permissions for package params
func (o *GetPermissionsForPackageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPermissionsForPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
