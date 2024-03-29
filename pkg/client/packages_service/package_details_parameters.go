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

// NewPackageDetailsParams creates a new PackageDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackageDetailsParams() *PackageDetailsParams {
	return &PackageDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackageDetailsParamsWithTimeout creates a new PackageDetailsParams object
// with the ability to set a timeout on a request.
func NewPackageDetailsParamsWithTimeout(timeout time.Duration) *PackageDetailsParams {
	return &PackageDetailsParams{
		timeout: timeout,
	}
}

// NewPackageDetailsParamsWithContext creates a new PackageDetailsParams object
// with the ability to set a context for a request.
func NewPackageDetailsParamsWithContext(ctx context.Context) *PackageDetailsParams {
	return &PackageDetailsParams{
		Context: ctx,
	}
}

// NewPackageDetailsParamsWithHTTPClient creates a new PackageDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackageDetailsParamsWithHTTPClient(client *http.Client) *PackageDetailsParams {
	return &PackageDetailsParams{
		HTTPClient: client,
	}
}

/*
PackageDetailsParams contains all the parameters to send to the API endpoint

	for the package details operation.

	Typically these are written to a http.Request.
*/
type PackageDetailsParams struct {

	// PackageName.
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the package details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageDetailsParams) WithDefaults() *PackageDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the package details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the package details params
func (o *PackageDetailsParams) WithTimeout(timeout time.Duration) *PackageDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package details params
func (o *PackageDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package details params
func (o *PackageDetailsParams) WithContext(ctx context.Context) *PackageDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package details params
func (o *PackageDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package details params
func (o *PackageDetailsParams) WithHTTPClient(client *http.Client) *PackageDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package details params
func (o *PackageDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageName adds the packageName to the package details params
func (o *PackageDetailsParams) WithPackageName(packageName string) *PackageDetailsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the package details params
func (o *PackageDetailsParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *PackageDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
