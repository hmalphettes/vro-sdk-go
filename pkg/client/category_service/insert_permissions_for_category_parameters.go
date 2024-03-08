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

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// NewInsertPermissionsForCategoryParams creates a new InsertPermissionsForCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInsertPermissionsForCategoryParams() *InsertPermissionsForCategoryParams {
	return &InsertPermissionsForCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInsertPermissionsForCategoryParamsWithTimeout creates a new InsertPermissionsForCategoryParams object
// with the ability to set a timeout on a request.
func NewInsertPermissionsForCategoryParamsWithTimeout(timeout time.Duration) *InsertPermissionsForCategoryParams {
	return &InsertPermissionsForCategoryParams{
		timeout: timeout,
	}
}

// NewInsertPermissionsForCategoryParamsWithContext creates a new InsertPermissionsForCategoryParams object
// with the ability to set a context for a request.
func NewInsertPermissionsForCategoryParamsWithContext(ctx context.Context) *InsertPermissionsForCategoryParams {
	return &InsertPermissionsForCategoryParams{
		Context: ctx,
	}
}

// NewInsertPermissionsForCategoryParamsWithHTTPClient creates a new InsertPermissionsForCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewInsertPermissionsForCategoryParamsWithHTTPClient(client *http.Client) *InsertPermissionsForCategoryParams {
	return &InsertPermissionsForCategoryParams{
		HTTPClient: client,
	}
}

/*
InsertPermissionsForCategoryParams contains all the parameters to send to the API endpoint

	for the insert permissions for category operation.

	Typically these are written to a http.Request.
*/
type InsertPermissionsForCategoryParams struct {

	// Body.
	Body *models.Permissions

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the insert permissions for category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertPermissionsForCategoryParams) WithDefaults() *InsertPermissionsForCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the insert permissions for category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertPermissionsForCategoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) WithTimeout(timeout time.Duration) *InsertPermissionsForCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) WithContext(ctx context.Context) *InsertPermissionsForCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) WithHTTPClient(client *http.Client) *InsertPermissionsForCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) WithBody(body *models.Permissions) *InsertPermissionsForCategoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) SetBody(body *models.Permissions) {
	o.Body = body
}

// WithID adds the id to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) WithID(id string) *InsertPermissionsForCategoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the insert permissions for category params
func (o *InsertPermissionsForCategoryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InsertPermissionsForCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
