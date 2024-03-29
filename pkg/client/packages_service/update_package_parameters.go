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

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// NewUpdatePackageParams creates a new UpdatePackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePackageParams() *UpdatePackageParams {
	return &UpdatePackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePackageParamsWithTimeout creates a new UpdatePackageParams object
// with the ability to set a timeout on a request.
func NewUpdatePackageParamsWithTimeout(timeout time.Duration) *UpdatePackageParams {
	return &UpdatePackageParams{
		timeout: timeout,
	}
}

// NewUpdatePackageParamsWithContext creates a new UpdatePackageParams object
// with the ability to set a context for a request.
func NewUpdatePackageParamsWithContext(ctx context.Context) *UpdatePackageParams {
	return &UpdatePackageParams{
		Context: ctx,
	}
}

// NewUpdatePackageParamsWithHTTPClient creates a new UpdatePackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePackageParamsWithHTTPClient(client *http.Client) *UpdatePackageParams {
	return &UpdatePackageParams{
		HTTPClient: client,
	}
}

/*
UpdatePackageParams contains all the parameters to send to the API endpoint

	for the update package operation.

	Typically these are written to a http.Request.
*/
type UpdatePackageParams struct {

	// XVROChangesetSha.
	XVROChangesetSha *string

	// Body.
	Body *models.UpdatePackageDetails

	// PackageName.
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePackageParams) WithDefaults() *UpdatePackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePackageParams) SetDefaults() {
	var (
		xVROChangesetShaDefault = string("")
	)

	val := UpdatePackageParams{
		XVROChangesetSha: &xVROChangesetShaDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update package params
func (o *UpdatePackageParams) WithTimeout(timeout time.Duration) *UpdatePackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update package params
func (o *UpdatePackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update package params
func (o *UpdatePackageParams) WithContext(ctx context.Context) *UpdatePackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update package params
func (o *UpdatePackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update package params
func (o *UpdatePackageParams) WithHTTPClient(client *http.Client) *UpdatePackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update package params
func (o *UpdatePackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVROChangesetSha adds the xVROChangesetSha to the update package params
func (o *UpdatePackageParams) WithXVROChangesetSha(xVROChangesetSha *string) *UpdatePackageParams {
	o.SetXVROChangesetSha(xVROChangesetSha)
	return o
}

// SetXVROChangesetSha adds the xVROChangesetSha to the update package params
func (o *UpdatePackageParams) SetXVROChangesetSha(xVROChangesetSha *string) {
	o.XVROChangesetSha = xVROChangesetSha
}

// WithBody adds the body to the update package params
func (o *UpdatePackageParams) WithBody(body *models.UpdatePackageDetails) *UpdatePackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update package params
func (o *UpdatePackageParams) SetBody(body *models.UpdatePackageDetails) {
	o.Body = body
}

// WithPackageName adds the packageName to the update package params
func (o *UpdatePackageParams) WithPackageName(packageName string) *UpdatePackageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the update package params
func (o *UpdatePackageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XVROChangesetSha != nil {

		// header param X-VRO-Changeset-Sha
		if err := r.SetHeaderParam("X-VRO-Changeset-Sha", *o.XVROChangesetSha); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
