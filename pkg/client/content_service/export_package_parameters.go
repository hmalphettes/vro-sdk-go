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

// NewExportPackageParams creates a new ExportPackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportPackageParams() *ExportPackageParams {
	return &ExportPackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportPackageParamsWithTimeout creates a new ExportPackageParams object
// with the ability to set a timeout on a request.
func NewExportPackageParamsWithTimeout(timeout time.Duration) *ExportPackageParams {
	return &ExportPackageParams{
		timeout: timeout,
	}
}

// NewExportPackageParamsWithContext creates a new ExportPackageParams object
// with the ability to set a context for a request.
func NewExportPackageParamsWithContext(ctx context.Context) *ExportPackageParams {
	return &ExportPackageParams{
		Context: ctx,
	}
}

// NewExportPackageParamsWithHTTPClient creates a new ExportPackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportPackageParamsWithHTTPClient(client *http.Client) *ExportPackageParams {
	return &ExportPackageParams{
		HTTPClient: client,
	}
}

/*
ExportPackageParams contains all the parameters to send to the API endpoint

	for the export package operation.

	Typically these are written to a http.Request.
*/
type ExportPackageParams struct {

	// ExportConfigSecureStringAttributeValues.
	ExportConfigSecureStringAttributeValues *bool

	// ExportConfigurationAttributeValues.
	//
	// Default: true
	ExportConfigurationAttributeValues *bool

	// ExportGlobalTags.
	//
	// Default: true
	ExportGlobalTags *bool

	// ExportVersionHistory.
	//
	// Default: true
	ExportVersionHistory *bool

	// PackageName.
	PackageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportPackageParams) WithDefaults() *ExportPackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportPackageParams) SetDefaults() {
	var (
		exportConfigSecureStringAttributeValuesDefault = bool(false)

		exportConfigurationAttributeValuesDefault = bool(true)

		exportGlobalTagsDefault = bool(true)

		exportVersionHistoryDefault = bool(true)
	)

	val := ExportPackageParams{
		ExportConfigSecureStringAttributeValues: &exportConfigSecureStringAttributeValuesDefault,
		ExportConfigurationAttributeValues:      &exportConfigurationAttributeValuesDefault,
		ExportGlobalTags:                        &exportGlobalTagsDefault,
		ExportVersionHistory:                    &exportVersionHistoryDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the export package params
func (o *ExportPackageParams) WithTimeout(timeout time.Duration) *ExportPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export package params
func (o *ExportPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export package params
func (o *ExportPackageParams) WithContext(ctx context.Context) *ExportPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export package params
func (o *ExportPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export package params
func (o *ExportPackageParams) WithHTTPClient(client *http.Client) *ExportPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export package params
func (o *ExportPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExportConfigSecureStringAttributeValues adds the exportConfigSecureStringAttributeValues to the export package params
func (o *ExportPackageParams) WithExportConfigSecureStringAttributeValues(exportConfigSecureStringAttributeValues *bool) *ExportPackageParams {
	o.SetExportConfigSecureStringAttributeValues(exportConfigSecureStringAttributeValues)
	return o
}

// SetExportConfigSecureStringAttributeValues adds the exportConfigSecureStringAttributeValues to the export package params
func (o *ExportPackageParams) SetExportConfigSecureStringAttributeValues(exportConfigSecureStringAttributeValues *bool) {
	o.ExportConfigSecureStringAttributeValues = exportConfigSecureStringAttributeValues
}

// WithExportConfigurationAttributeValues adds the exportConfigurationAttributeValues to the export package params
func (o *ExportPackageParams) WithExportConfigurationAttributeValues(exportConfigurationAttributeValues *bool) *ExportPackageParams {
	o.SetExportConfigurationAttributeValues(exportConfigurationAttributeValues)
	return o
}

// SetExportConfigurationAttributeValues adds the exportConfigurationAttributeValues to the export package params
func (o *ExportPackageParams) SetExportConfigurationAttributeValues(exportConfigurationAttributeValues *bool) {
	o.ExportConfigurationAttributeValues = exportConfigurationAttributeValues
}

// WithExportGlobalTags adds the exportGlobalTags to the export package params
func (o *ExportPackageParams) WithExportGlobalTags(exportGlobalTags *bool) *ExportPackageParams {
	o.SetExportGlobalTags(exportGlobalTags)
	return o
}

// SetExportGlobalTags adds the exportGlobalTags to the export package params
func (o *ExportPackageParams) SetExportGlobalTags(exportGlobalTags *bool) {
	o.ExportGlobalTags = exportGlobalTags
}

// WithExportVersionHistory adds the exportVersionHistory to the export package params
func (o *ExportPackageParams) WithExportVersionHistory(exportVersionHistory *bool) *ExportPackageParams {
	o.SetExportVersionHistory(exportVersionHistory)
	return o
}

// SetExportVersionHistory adds the exportVersionHistory to the export package params
func (o *ExportPackageParams) SetExportVersionHistory(exportVersionHistory *bool) {
	o.ExportVersionHistory = exportVersionHistory
}

// WithPackageName adds the packageName to the export package params
func (o *ExportPackageParams) WithPackageName(packageName string) *ExportPackageParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the export package params
func (o *ExportPackageParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WriteToRequest writes these params to a swagger request
func (o *ExportPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExportConfigSecureStringAttributeValues != nil {

		// query param exportConfigSecureStringAttributeValues
		var qrExportConfigSecureStringAttributeValues bool

		if o.ExportConfigSecureStringAttributeValues != nil {
			qrExportConfigSecureStringAttributeValues = *o.ExportConfigSecureStringAttributeValues
		}
		qExportConfigSecureStringAttributeValues := swag.FormatBool(qrExportConfigSecureStringAttributeValues)
		if qExportConfigSecureStringAttributeValues != "" {

			if err := r.SetQueryParam("exportConfigSecureStringAttributeValues", qExportConfigSecureStringAttributeValues); err != nil {
				return err
			}
		}
	}

	if o.ExportConfigurationAttributeValues != nil {

		// query param exportConfigurationAttributeValues
		var qrExportConfigurationAttributeValues bool

		if o.ExportConfigurationAttributeValues != nil {
			qrExportConfigurationAttributeValues = *o.ExportConfigurationAttributeValues
		}
		qExportConfigurationAttributeValues := swag.FormatBool(qrExportConfigurationAttributeValues)
		if qExportConfigurationAttributeValues != "" {

			if err := r.SetQueryParam("exportConfigurationAttributeValues", qExportConfigurationAttributeValues); err != nil {
				return err
			}
		}
	}

	if o.ExportGlobalTags != nil {

		// query param exportGlobalTags
		var qrExportGlobalTags bool

		if o.ExportGlobalTags != nil {
			qrExportGlobalTags = *o.ExportGlobalTags
		}
		qExportGlobalTags := swag.FormatBool(qrExportGlobalTags)
		if qExportGlobalTags != "" {

			if err := r.SetQueryParam("exportGlobalTags", qExportGlobalTags); err != nil {
				return err
			}
		}
	}

	if o.ExportVersionHistory != nil {

		// query param exportVersionHistory
		var qrExportVersionHistory bool

		if o.ExportVersionHistory != nil {
			qrExportVersionHistory = *o.ExportVersionHistory
		}
		qExportVersionHistory := swag.FormatBool(qrExportVersionHistory)
		if qExportVersionHistory != "" {

			if err := r.SetQueryParam("exportVersionHistory", qExportVersionHistory); err != nil {
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
