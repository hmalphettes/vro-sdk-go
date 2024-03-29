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
	"github.com/go-openapi/swag"
)

// NewDeleteContentPackageParams creates a new DeleteContentPackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContentPackageParams() *DeleteContentPackageParams {
	return &DeleteContentPackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentPackageParamsWithTimeout creates a new DeleteContentPackageParams object
// with the ability to set a timeout on a request.
func NewDeleteContentPackageParamsWithTimeout(timeout time.Duration) *DeleteContentPackageParams {
	return &DeleteContentPackageParams{
		timeout: timeout,
	}
}

// NewDeleteContentPackageParamsWithContext creates a new DeleteContentPackageParams object
// with the ability to set a context for a request.
func NewDeleteContentPackageParamsWithContext(ctx context.Context) *DeleteContentPackageParams {
	return &DeleteContentPackageParams{
		Context: ctx,
	}
}

// NewDeleteContentPackageParamsWithHTTPClient creates a new DeleteContentPackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContentPackageParamsWithHTTPClient(client *http.Client) *DeleteContentPackageParams {
	return &DeleteContentPackageParams{
		HTTPClient: client,
	}
}

/*
DeleteContentPackageParams contains all the parameters to send to the API endpoint

	for the delete content package operation.

	Typically these are written to a http.Request.
*/
type DeleteContentPackageParams struct {

	// Force.
	Force *bool

	// PackageName.
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete content package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentPackageParams) WithDefaults() *DeleteContentPackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete content package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentPackageParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := DeleteContentPackageParams{
		Force: &forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete content package params
func (o *DeleteContentPackageParams) WithTimeout(timeout time.Duration) *DeleteContentPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete content package params
func (o *DeleteContentPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete content package params
func (o *DeleteContentPackageParams) WithContext(ctx context.Context) *DeleteContentPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete content package params
func (o *DeleteContentPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete content package params
func (o *DeleteContentPackageParams) WithHTTPClient(client *http.Client) *DeleteContentPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete content package params
func (o *DeleteContentPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the delete content package params
func (o *DeleteContentPackageParams) WithForce(force *bool) *DeleteContentPackageParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete content package params
func (o *DeleteContentPackageParams) SetForce(force *bool) {
	o.Force = force
}

// WithPackageName adds the packageName to the delete content package params
func (o *DeleteContentPackageParams) WithPackageName(packageName string) *DeleteContentPackageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the delete content package params
func (o *DeleteContentPackageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
