// Code generated by go-swagger; DO NOT EDIT.

package catalog_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DownloadIconForModule(params *DownloadIconForModuleParams, opts ...ClientOption) (*DownloadIconForModuleOK, error)

	DownloadIconForType(params *DownloadIconForTypeParams, opts ...ClientOption) (*DownloadIconForTypeOK, error)

	FetchPluginMetadata(params *FetchPluginMetadataParams, opts ...ClientOption) (*FetchPluginMetadataOK, error)

	FindByID(params *FindByIDParams, opts ...ClientOption) (*FindByIDOK, error)

	FindByRelation(params *FindByRelationParams, opts ...ClientOption) (*FindByRelationOK, error)

	FindRootElement(params *FindRootElementParams, opts ...ClientOption) (*FindRootElementOK, error)

	FindSimpleListQuery(params *FindSimpleListQueryParams, opts ...ClientOption) (*FindSimpleListQueryOK, error)

	ListNamespaces(params *ListNamespacesParams, opts ...ClientOption) (*ListNamespacesOK, error)

	ListObjectTags(params *ListObjectTagsParams, opts ...ClientOption) (*ListObjectTagsOK, error)

	TagObject(params *TagObjectParams, opts ...ClientOption) (*TagObjectNoContent, error)

	UntagObject(params *UntagObjectParams, opts ...ClientOption) (*UntagObjectNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DownloadIconForModule downloads icon for module

If the request is successful, the API responds with an HTTP 200 OK status code and the requested icon image as a downloadable attachment that has an 'image/png' MIME type.
*/
func (a *Client) DownloadIconForModule(params *DownloadIconForModuleParams, opts ...ClientOption) (*DownloadIconForModuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadIconForModuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadIconForModule",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}/metadata/icon",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadIconForModuleReader{formats: a.formats},
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
	success, ok := result.(*DownloadIconForModuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadIconForModule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DownloadIconForType downloads icon for type

All types that a Orchestrator plug-in defines, and the system types that the Orchestrator server defines, have default icon images. If the request is successful, the API responds with an HTTP 200 OK status code and the requested icon image as a downloadable attachment that has an 'image/png' MIME type.
*/
func (a *Client) DownloadIconForType(params *DownloadIconForTypeParams, opts ...ClientOption) (*DownloadIconForTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadIconForTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadIconForType",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}/{type}/metadata/icon",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadIconForTypeReader{formats: a.formats},
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
	success, ok := result.(*DownloadIconForTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadIconForType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FetchPluginMetadata fetches plugin metadata

Plug-in metadata contains information about the types that this plug-in defines and their attributes or relations.
*/
func (a *Client) FetchPluginMetadata(params *FetchPluginMetadataParams, opts ...ClientOption) (*FetchPluginMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchPluginMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "fetchPluginMetadata",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}/metadata",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchPluginMetadataReader{formats: a.formats},
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
	success, ok := result.(*FetchPluginMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetchPluginMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindByID finds by id

A REST endpoint for retrieving an element by its type and ID. All Orchestrator objects can be accessed or identified by their unique combination of namespace, type and ID.
*/
func (a *Client) FindByID(params *FindByIDParams, opts ...ClientOption) (*FindByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findById",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}/{type}/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindByIDReader{formats: a.formats},
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
	success, ok := result.(*FindByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindByRelation finds by relation

Every Orchestrator plug-in can define relations between its types. For example, hierarchy structures can be defined as 'CHILDREN' relations between parent and children types. If the request is successful, the API responds with an HTTP 200 OK status code and a list containing the child objects that are accessible through a given relation, if any.
*/
func (a *Client) FindByRelation(params *FindByRelationParams, opts ...ClientOption) (*FindByRelationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindByRelationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findByRelation",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}/{parentType}/{parentId}/{relationName}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindByRelationReader{formats: a.formats},
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
	success, ok := result.(*FindByRelationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findByRelation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindRootElement finds root element

A REST endpoint for retrieving the root elements in a namespace. All Orchestrator plug-ins that have an inventory, usually provide a single root element. The Orchestrator server does not have a single root. Instead, for the System namespace, the Orchestrator REST API returns a list that contains links to all system types (workflows, tasks, etc.)
*/
func (a *Client) FindRootElement(params *FindRootElementParams, opts ...ClientOption) (*FindRootElementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindRootElementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findRootElement",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindRootElementReader{formats: a.formats},
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
	success, ok := result.(*FindRootElementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findRootElement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindSimpleListQuery lists objects for specific type

If the request is successful, the API responds with an HTTP 200 OK status code and the requested list of objects
*/
func (a *Client) FindSimpleListQuery(params *FindSimpleListQueryParams, opts ...ClientOption) (*FindSimpleListQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSimpleListQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSimpleListQuery",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}/{type}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindSimpleListQueryReader{formats: a.formats},
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
	success, ok := result.(*FindSimpleListQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSimpleListQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListNamespaces lists namespaces

A REST endpoint for listing all catalog entry points. Each Orchestrator plug-in exposes its own catalog entry point that is accessible at /catalog/{pluginname}. The Orchestrator server exposes a separate catalog entry point that is accessible at /catalog/System
*/
func (a *Client) ListNamespaces(params *ListNamespacesParams, opts ...ClientOption) (*ListNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNamespacesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNamespaces",
		Method:             "GET",
		PathPattern:        "/catalog",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNamespacesReader{formats: a.formats},
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
	success, ok := result.(*ListNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListObjectTags lists object tags

A REST endpoint to retrieve tags attached to an element.
*/
func (a *Client) ListObjectTags(params *ListObjectTagsParams, opts ...ClientOption) (*ListObjectTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListObjectTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listObjectTags",
		Method:             "GET",
		PathPattern:        "/catalog/{namespace}/{type}/{id}/tags",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListObjectTagsReader{formats: a.formats},
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
	success, ok := result.(*ListObjectTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listObjectTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TagObject attaches tag to entity

A REST endpoint for attaching tag to entity.
*/
func (a *Client) TagObject(params *TagObjectParams, opts ...ClientOption) (*TagObjectNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTagObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "tagObject",
		Method:             "POST",
		PathPattern:        "/catalog/{namespace}/{type}/{id}/tags",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TagObjectReader{formats: a.formats},
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
	success, ok := result.(*TagObjectNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tagObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UntagObject removes tag from entity

To remove global tag form entity tag name must be prefixed with ":".
*/
func (a *Client) UntagObject(params *UntagObjectParams, opts ...ClientOption) (*UntagObjectNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUntagObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "untagObject",
		Method:             "DELETE",
		PathPattern:        "/catalog/{namespace}/{type}/{id}/tags/{tagname}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UntagObjectReader{formats: a.formats},
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
	success, ok := result.(*UntagObjectNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for untagObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
