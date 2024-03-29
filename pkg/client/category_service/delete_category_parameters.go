// Code generated by go-swagger; DO NOT EDIT.

package category_service

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

// NewDeleteCategoryParams creates a new DeleteCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCategoryParams() *DeleteCategoryParams {
	return &DeleteCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCategoryParamsWithTimeout creates a new DeleteCategoryParams object
// with the ability to set a timeout on a request.
func NewDeleteCategoryParamsWithTimeout(timeout time.Duration) *DeleteCategoryParams {
	return &DeleteCategoryParams{
		timeout: timeout,
	}
}

// NewDeleteCategoryParamsWithContext creates a new DeleteCategoryParams object
// with the ability to set a context for a request.
func NewDeleteCategoryParamsWithContext(ctx context.Context) *DeleteCategoryParams {
	return &DeleteCategoryParams{
		Context: ctx,
	}
}

// NewDeleteCategoryParamsWithHTTPClient creates a new DeleteCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCategoryParamsWithHTTPClient(client *http.Client) *DeleteCategoryParams {
	return &DeleteCategoryParams{
		HTTPClient: client,
	}
}

/*
DeleteCategoryParams contains all the parameters to send to the API endpoint

	for the delete category operation.

	Typically these are written to a http.Request.
*/
type DeleteCategoryParams struct {

	// DeleteNonEmptyContent.
	DeleteNonEmptyContent *bool

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCategoryParams) WithDefaults() *DeleteCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCategoryParams) SetDefaults() {
	var (
		deleteNonEmptyContentDefault = bool(false)
	)

	val := DeleteCategoryParams{
		DeleteNonEmptyContent: &deleteNonEmptyContentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete category params
func (o *DeleteCategoryParams) WithTimeout(timeout time.Duration) *DeleteCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete category params
func (o *DeleteCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete category params
func (o *DeleteCategoryParams) WithContext(ctx context.Context) *DeleteCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete category params
func (o *DeleteCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete category params
func (o *DeleteCategoryParams) WithHTTPClient(client *http.Client) *DeleteCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete category params
func (o *DeleteCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteNonEmptyContent adds the deleteNonEmptyContent to the delete category params
func (o *DeleteCategoryParams) WithDeleteNonEmptyContent(deleteNonEmptyContent *bool) *DeleteCategoryParams {
	o.SetDeleteNonEmptyContent(deleteNonEmptyContent)
	return o
}

// SetDeleteNonEmptyContent adds the deleteNonEmptyContent to the delete category params
func (o *DeleteCategoryParams) SetDeleteNonEmptyContent(deleteNonEmptyContent *bool) {
	o.DeleteNonEmptyContent = deleteNonEmptyContent
}

// WithID adds the id to the delete category params
func (o *DeleteCategoryParams) WithID(id string) *DeleteCategoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete category params
func (o *DeleteCategoryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteNonEmptyContent != nil {

		// query param deleteNonEmptyContent
		var qrDeleteNonEmptyContent bool

		if o.DeleteNonEmptyContent != nil {
			qrDeleteNonEmptyContent = *o.DeleteNonEmptyContent
		}
		qDeleteNonEmptyContent := swag.FormatBool(qrDeleteNonEmptyContent)
		if qDeleteNonEmptyContent != "" {

			if err := r.SetQueryParam("deleteNonEmptyContent", qDeleteNonEmptyContent); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
