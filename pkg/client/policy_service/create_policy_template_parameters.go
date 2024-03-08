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

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// NewCreatePolicyTemplateParams creates a new CreatePolicyTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePolicyTemplateParams() *CreatePolicyTemplateParams {
	return &CreatePolicyTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyTemplateParamsWithTimeout creates a new CreatePolicyTemplateParams object
// with the ability to set a timeout on a request.
func NewCreatePolicyTemplateParamsWithTimeout(timeout time.Duration) *CreatePolicyTemplateParams {
	return &CreatePolicyTemplateParams{
		timeout: timeout,
	}
}

// NewCreatePolicyTemplateParamsWithContext creates a new CreatePolicyTemplateParams object
// with the ability to set a context for a request.
func NewCreatePolicyTemplateParamsWithContext(ctx context.Context) *CreatePolicyTemplateParams {
	return &CreatePolicyTemplateParams{
		Context: ctx,
	}
}

// NewCreatePolicyTemplateParamsWithHTTPClient creates a new CreatePolicyTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePolicyTemplateParamsWithHTTPClient(client *http.Client) *CreatePolicyTemplateParams {
	return &CreatePolicyTemplateParams{
		HTTPClient: client,
	}
}

/*
CreatePolicyTemplateParams contains all the parameters to send to the API endpoint

	for the create policy template operation.

	Typically these are written to a http.Request.
*/
type CreatePolicyTemplateParams struct {

	// Body.
	Body *models.PolicyTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create policy template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyTemplateParams) WithDefaults() *CreatePolicyTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create policy template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create policy template params
func (o *CreatePolicyTemplateParams) WithTimeout(timeout time.Duration) *CreatePolicyTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy template params
func (o *CreatePolicyTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy template params
func (o *CreatePolicyTemplateParams) WithContext(ctx context.Context) *CreatePolicyTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy template params
func (o *CreatePolicyTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy template params
func (o *CreatePolicyTemplateParams) WithHTTPClient(client *http.Client) *CreatePolicyTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy template params
func (o *CreatePolicyTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create policy template params
func (o *CreatePolicyTemplateParams) WithBody(body *models.PolicyTemplate) *CreatePolicyTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create policy template params
func (o *CreatePolicyTemplateParams) SetBody(body *models.PolicyTemplate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
