// Code generated by go-swagger; DO NOT EDIT.

package authorization_groups_service

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateAuthorizationGroupParams creates a new CreateAuthorizationGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAuthorizationGroupParams() *CreateAuthorizationGroupParams {
	return &CreateAuthorizationGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthorizationGroupParamsWithTimeout creates a new CreateAuthorizationGroupParams object
// with the ability to set a timeout on a request.
func NewCreateAuthorizationGroupParamsWithTimeout(timeout time.Duration) *CreateAuthorizationGroupParams {
	return &CreateAuthorizationGroupParams{
		timeout: timeout,
	}
}

// NewCreateAuthorizationGroupParamsWithContext creates a new CreateAuthorizationGroupParams object
// with the ability to set a context for a request.
func NewCreateAuthorizationGroupParamsWithContext(ctx context.Context) *CreateAuthorizationGroupParams {
	return &CreateAuthorizationGroupParams{
		Context: ctx,
	}
}

// NewCreateAuthorizationGroupParamsWithHTTPClient creates a new CreateAuthorizationGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAuthorizationGroupParamsWithHTTPClient(client *http.Client) *CreateAuthorizationGroupParams {
	return &CreateAuthorizationGroupParams{
		HTTPClient: client,
	}
}

/*
CreateAuthorizationGroupParams contains all the parameters to send to the API endpoint

	for the create authorization group operation.

	Typically these are written to a http.Request.
*/
type CreateAuthorizationGroupParams struct {

	// Body.
	Body *models.AuthorizationGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create authorization group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthorizationGroupParams) WithDefaults() *CreateAuthorizationGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create authorization group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthorizationGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create authorization group params
func (o *CreateAuthorizationGroupParams) WithTimeout(timeout time.Duration) *CreateAuthorizationGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create authorization group params
func (o *CreateAuthorizationGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create authorization group params
func (o *CreateAuthorizationGroupParams) WithContext(ctx context.Context) *CreateAuthorizationGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create authorization group params
func (o *CreateAuthorizationGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create authorization group params
func (o *CreateAuthorizationGroupParams) WithHTTPClient(client *http.Client) *CreateAuthorizationGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create authorization group params
func (o *CreateAuthorizationGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create authorization group params
func (o *CreateAuthorizationGroupParams) WithBody(body *models.AuthorizationGroup) *CreateAuthorizationGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create authorization group params
func (o *CreateAuthorizationGroupParams) SetBody(body *models.AuthorizationGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthorizationGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
