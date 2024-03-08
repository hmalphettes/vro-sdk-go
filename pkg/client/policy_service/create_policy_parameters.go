// Code generated by go-swagger; DO NOT EDIT.

package policy_service

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

// NewCreatePolicyParams creates a new CreatePolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePolicyParams() *CreatePolicyParams {
	return &CreatePolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyParamsWithTimeout creates a new CreatePolicyParams object
// with the ability to set a timeout on a request.
func NewCreatePolicyParamsWithTimeout(timeout time.Duration) *CreatePolicyParams {
	return &CreatePolicyParams{
		timeout: timeout,
	}
}

// NewCreatePolicyParamsWithContext creates a new CreatePolicyParams object
// with the ability to set a context for a request.
func NewCreatePolicyParamsWithContext(ctx context.Context) *CreatePolicyParams {
	return &CreatePolicyParams{
		Context: ctx,
	}
}

// NewCreatePolicyParamsWithHTTPClient creates a new CreatePolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePolicyParamsWithHTTPClient(client *http.Client) *CreatePolicyParams {
	return &CreatePolicyParams{
		HTTPClient: client,
	}
}

/*
CreatePolicyParams contains all the parameters to send to the API endpoint

	for the create policy operation.

	Typically these are written to a http.Request.
*/
type CreatePolicyParams struct {

	// Body.
	Body *models.Policy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyParams) WithDefaults() *CreatePolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create policy params
func (o *CreatePolicyParams) WithTimeout(timeout time.Duration) *CreatePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy params
func (o *CreatePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy params
func (o *CreatePolicyParams) WithContext(ctx context.Context) *CreatePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy params
func (o *CreatePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy params
func (o *CreatePolicyParams) WithHTTPClient(client *http.Client) *CreatePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy params
func (o *CreatePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create policy params
func (o *CreatePolicyParams) WithBody(body *models.Policy) *CreatePolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create policy params
func (o *CreatePolicyParams) SetBody(body *models.Policy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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