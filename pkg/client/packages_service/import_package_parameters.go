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
	"github.com/go-openapi/swag"
)

// NewImportPackageParams creates a new ImportPackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportPackageParams() *ImportPackageParams {
	return &ImportPackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportPackageParamsWithTimeout creates a new ImportPackageParams object
// with the ability to set a timeout on a request.
func NewImportPackageParamsWithTimeout(timeout time.Duration) *ImportPackageParams {
	return &ImportPackageParams{
		timeout: timeout,
	}
}

// NewImportPackageParamsWithContext creates a new ImportPackageParams object
// with the ability to set a context for a request.
func NewImportPackageParamsWithContext(ctx context.Context) *ImportPackageParams {
	return &ImportPackageParams{
		Context: ctx,
	}
}

// NewImportPackageParamsWithHTTPClient creates a new ImportPackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportPackageParamsWithHTTPClient(client *http.Client) *ImportPackageParams {
	return &ImportPackageParams{
		HTTPClient: client,
	}
}

/*
ImportPackageParams contains all the parameters to send to the API endpoint

	for the import package operation.

	Typically these are written to a http.Request.
*/
type ImportPackageParams struct {

	// File.
	File runtime.NamedReadCloser

	// ImportConfigSecureStringAttributeValues.
	ImportConfigSecureStringAttributeValues *bool

	// ImportConfigurationAttributeValues.
	//
	// Default: true
	ImportConfigurationAttributeValues *bool

	// Overwrite.
	Overwrite *bool

	// TagImportMode.
	//
	// Default: "ImportButPreserveExistingValue"
	TagImportMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportPackageParams) WithDefaults() *ImportPackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportPackageParams) SetDefaults() {
	var (
		importConfigSecureStringAttributeValuesDefault = bool(false)

		importConfigurationAttributeValuesDefault = bool(true)

		overwriteDefault = bool(false)

		tagImportModeDefault = string("ImportButPreserveExistingValue")
	)

	val := ImportPackageParams{
		ImportConfigSecureStringAttributeValues: &importConfigSecureStringAttributeValuesDefault,
		ImportConfigurationAttributeValues:      &importConfigurationAttributeValuesDefault,
		Overwrite:                               &overwriteDefault,
		TagImportMode:                           &tagImportModeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the import package params
func (o *ImportPackageParams) WithTimeout(timeout time.Duration) *ImportPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import package params
func (o *ImportPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import package params
func (o *ImportPackageParams) WithContext(ctx context.Context) *ImportPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import package params
func (o *ImportPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import package params
func (o *ImportPackageParams) WithHTTPClient(client *http.Client) *ImportPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import package params
func (o *ImportPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the import package params
func (o *ImportPackageParams) WithFile(file runtime.NamedReadCloser) *ImportPackageParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the import package params
func (o *ImportPackageParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithImportConfigSecureStringAttributeValues adds the importConfigSecureStringAttributeValues to the import package params
func (o *ImportPackageParams) WithImportConfigSecureStringAttributeValues(importConfigSecureStringAttributeValues *bool) *ImportPackageParams {
	o.SetImportConfigSecureStringAttributeValues(importConfigSecureStringAttributeValues)
	return o
}

// SetImportConfigSecureStringAttributeValues adds the importConfigSecureStringAttributeValues to the import package params
func (o *ImportPackageParams) SetImportConfigSecureStringAttributeValues(importConfigSecureStringAttributeValues *bool) {
	o.ImportConfigSecureStringAttributeValues = importConfigSecureStringAttributeValues
}

// WithImportConfigurationAttributeValues adds the importConfigurationAttributeValues to the import package params
func (o *ImportPackageParams) WithImportConfigurationAttributeValues(importConfigurationAttributeValues *bool) *ImportPackageParams {
	o.SetImportConfigurationAttributeValues(importConfigurationAttributeValues)
	return o
}

// SetImportConfigurationAttributeValues adds the importConfigurationAttributeValues to the import package params
func (o *ImportPackageParams) SetImportConfigurationAttributeValues(importConfigurationAttributeValues *bool) {
	o.ImportConfigurationAttributeValues = importConfigurationAttributeValues
}

// WithOverwrite adds the overwrite to the import package params
func (o *ImportPackageParams) WithOverwrite(overwrite *bool) *ImportPackageParams {
	o.SetOverwrite(overwrite)
	return o
}

// SetOverwrite adds the overwrite to the import package params
func (o *ImportPackageParams) SetOverwrite(overwrite *bool) {
	o.Overwrite = overwrite
}

// WithTagImportMode adds the tagImportMode to the import package params
func (o *ImportPackageParams) WithTagImportMode(tagImportMode *string) *ImportPackageParams {
	o.SetTagImportMode(tagImportMode)
	return o
}

// SetTagImportMode adds the tagImportMode to the import package params
func (o *ImportPackageParams) SetTagImportMode(tagImportMode *string) {
	o.TagImportMode = tagImportMode
}

// WriteToRequest writes these params to a swagger request
func (o *ImportPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if o.ImportConfigSecureStringAttributeValues != nil {

		// query param importConfigSecureStringAttributeValues
		var qrImportConfigSecureStringAttributeValues bool

		if o.ImportConfigSecureStringAttributeValues != nil {
			qrImportConfigSecureStringAttributeValues = *o.ImportConfigSecureStringAttributeValues
		}
		qImportConfigSecureStringAttributeValues := swag.FormatBool(qrImportConfigSecureStringAttributeValues)
		if qImportConfigSecureStringAttributeValues != "" {

			if err := r.SetQueryParam("importConfigSecureStringAttributeValues", qImportConfigSecureStringAttributeValues); err != nil {
				return err
			}
		}
	}

	if o.ImportConfigurationAttributeValues != nil {

		// query param importConfigurationAttributeValues
		var qrImportConfigurationAttributeValues bool

		if o.ImportConfigurationAttributeValues != nil {
			qrImportConfigurationAttributeValues = *o.ImportConfigurationAttributeValues
		}
		qImportConfigurationAttributeValues := swag.FormatBool(qrImportConfigurationAttributeValues)
		if qImportConfigurationAttributeValues != "" {

			if err := r.SetQueryParam("importConfigurationAttributeValues", qImportConfigurationAttributeValues); err != nil {
				return err
			}
		}
	}

	if o.Overwrite != nil {

		// query param overwrite
		var qrOverwrite bool

		if o.Overwrite != nil {
			qrOverwrite = *o.Overwrite
		}
		qOverwrite := swag.FormatBool(qrOverwrite)
		if qOverwrite != "" {

			if err := r.SetQueryParam("overwrite", qOverwrite); err != nil {
				return err
			}
		}
	}

	if o.TagImportMode != nil {

		// query param tagImportMode
		var qrTagImportMode string

		if o.TagImportMode != nil {
			qrTagImportMode = *o.TagImportMode
		}
		qTagImportMode := qrTagImportMode
		if qTagImportMode != "" {

			if err := r.SetQueryParam("tagImportMode", qTagImportMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
