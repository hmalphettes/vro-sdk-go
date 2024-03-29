// Code generated by go-swagger; DO NOT EDIT.

package service_descriptor_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service descriptor service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service descriptor service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AboutInfo(params *AboutInfoParams, opts ...ClientOption) (*AboutInfoOK, error)

	Docs(params *DocsParams, opts ...ClientOption) (*DocsOK, error)

	EnumerateServices(params *EnumerateServicesParams, opts ...ClientOption) (*EnumerateServicesOK, error)

	GetServiceDescriptorXMLSchema(params *GetServiceDescriptorXMLSchemaParams, opts ...ClientOption) (*GetServiceDescriptorXMLSchemaOK, error)

	GetServiceDescriptorsXMLSchema(params *GetServiceDescriptorsXMLSchemaParams, opts ...ClientOption) (*GetServiceDescriptorsXMLSchemaOK, error)

	HealthStatus(params *HealthStatusParams, opts ...ClientOption) (*HealthStatusOK, error)

	Status(params *StatusParams, opts ...ClientOption) (*StatusOK, error)

	SupportedAPIVersions(params *SupportedAPIVersionsParams, opts ...ClientOption) (*SupportedAPIVersionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AboutInfo gets about info

A REST endpoint for retrieving server build properties
*/
func (a *Client) AboutInfo(params *AboutInfoParams, opts ...ClientOption) (*AboutInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAboutInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "aboutInfo",
		Method:             "GET",
		PathPattern:        "/about",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AboutInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AboutInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for aboutInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Docs redirects docs to docs index html

Redirect '/vco/api/docs' to '/vco/api/docs/index.html'
*/
func (a *Client) Docs(params *DocsParams, opts ...ClientOption) (*DocsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "docs",
		Method:             "GET",
		PathPattern:        "/docs",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DocsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for docs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EnumerateServices enumerates services

Lists the available top-level service entry points.
*/
func (a *Client) EnumerateServices(params *EnumerateServicesParams, opts ...ClientOption) (*EnumerateServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnumerateServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enumerateServices",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnumerateServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnumerateServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enumerateServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetServiceDescriptorXMLSchema gets r e s t x s d schema file

The XSD schema file defines the elements and types used by the REST service. You can use it to generate stub classes (in Java or other programming language)
*/
func (a *Client) GetServiceDescriptorXMLSchema(params *GetServiceDescriptorXMLSchemaParams, opts ...ClientOption) (*GetServiceDescriptorXMLSchemaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceDescriptorXMLSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServiceDescriptorXmlSchema",
		Method:             "GET",
		PathPattern:        "/schema/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceDescriptorXMLSchemaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceDescriptorXMLSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceDescriptorXmlSchema: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetServiceDescriptorsXMLSchema gets r e s t x s d schema file

The XSD schema file defines the elements and types used by the REST service. You can use it to generate stub classes (in Java or other programming language).
*/
func (a *Client) GetServiceDescriptorsXMLSchema(params *GetServiceDescriptorsXMLSchemaParams, opts ...ClientOption) (*GetServiceDescriptorsXMLSchemaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceDescriptorsXMLSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServiceDescriptorsXmlSchema",
		Method:             "GET",
		PathPattern:        "/schema",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceDescriptorsXMLSchemaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceDescriptorsXMLSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServiceDescriptorsXmlSchema: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
HealthStatus gets health status

A REST endpoint for retrieving server health status calculated on its health components (DB, authentication and etc.).
*/
func (a *Client) HealthStatus(params *HealthStatusParams, opts ...ClientOption) (*HealthStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "healthStatus",
		Method:             "GET",
		PathPattern:        "/healthstatus",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HealthStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HealthStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for healthStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Status gets v r a registration state

Returns the registration state of Orchestrator in vRA component registry.
*/
func (a *Client) Status(params *StatusParams, opts ...ClientOption) (*StatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "status",
		Method:             "GET",
		PathPattern:        "/status",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for status: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SupportedAPIVersions lists supported API versions

Currently, there is only one version.
*/
func (a *Client) SupportedAPIVersions(params *SupportedAPIVersionsParams, opts ...ClientOption) (*SupportedAPIVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSupportedAPIVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "supportedApiVersions",
		Method:             "GET",
		PathPattern:        "/versions",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SupportedAPIVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SupportedAPIVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for supportedApiVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
