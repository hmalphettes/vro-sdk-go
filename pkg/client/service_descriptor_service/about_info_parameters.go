// Code generated by go-swagger; DO NOT EDIT.

package service_descriptor_service

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

// NewAboutInfoParams creates a new AboutInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAboutInfoParams() *AboutInfoParams {
	return &AboutInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAboutInfoParamsWithTimeout creates a new AboutInfoParams object
// with the ability to set a timeout on a request.
func NewAboutInfoParamsWithTimeout(timeout time.Duration) *AboutInfoParams {
	return &AboutInfoParams{
		timeout: timeout,
	}
}

// NewAboutInfoParamsWithContext creates a new AboutInfoParams object
// with the ability to set a context for a request.
func NewAboutInfoParamsWithContext(ctx context.Context) *AboutInfoParams {
	return &AboutInfoParams{
		Context: ctx,
	}
}

// NewAboutInfoParamsWithHTTPClient creates a new AboutInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewAboutInfoParamsWithHTTPClient(client *http.Client) *AboutInfoParams {
	return &AboutInfoParams{
		HTTPClient: client,
	}
}

/*
AboutInfoParams contains all the parameters to send to the API endpoint

	for the about info operation.

	Typically these are written to a http.Request.
*/
type AboutInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the about info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AboutInfoParams) WithDefaults() *AboutInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the about info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AboutInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the about info params
func (o *AboutInfoParams) WithTimeout(timeout time.Duration) *AboutInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the about info params
func (o *AboutInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the about info params
func (o *AboutInfoParams) WithContext(ctx context.Context) *AboutInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the about info params
func (o *AboutInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the about info params
func (o *AboutInfoParams) WithHTTPClient(client *http.Client) *AboutInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the about info params
func (o *AboutInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AboutInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
