// Code generated by go-swagger; DO NOT EDIT.

package workflow_user_interaction_presentation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflow user interaction presentation service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow user interaction presentation service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllInteractionPresentations(params *AllInteractionPresentationsParams, opts ...ClientOption) (*AllInteractionPresentationsOK, error)

	DeleteWorkflowUserInteractionPresentationExecution(params *DeleteWorkflowUserInteractionPresentationExecutionParams, opts ...ClientOption) (*DeleteWorkflowUserInteractionPresentationExecutionNoContent, error)

	GetPresentationForInteraction(params *GetPresentationForInteractionParams, opts ...ClientOption) (*GetPresentationForInteractionOK, error)

	LoadExecutionForWorkflowUserInteractionPresentation(params *LoadExecutionForWorkflowUserInteractionPresentationParams, opts ...ClientOption) (*LoadExecutionForWorkflowUserInteractionPresentationOK, error)

	StartInteractionPresentation(params *StartInteractionPresentationParams, opts ...ClientOption) (*StartInteractionPresentationOK, error)

	UpdateWorkflowUserInteractionPresentation(params *UpdateWorkflowUserInteractionPresentationParams, opts ...ClientOption) (*UpdateWorkflowUserInteractionPresentationOK, error)

	UpdateWorkflowUserInteractionPresentationPUT(params *UpdateWorkflowUserInteractionPresentationPUTParams, opts ...ClientOption) (*UpdateWorkflowUserInteractionPresentationPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AllInteractionPresentations gets all interaction presentations

Retrieves a list of all user interaction presentations for a specified workflow execution. The returned list contains all currently running workflow execution user interaction presentation instances. If the user has admin rights, all presentation instances for all users are returned.
*/
func (a *Client) AllInteractionPresentations(params *AllInteractionPresentationsParams, opts ...ClientOption) (*AllInteractionPresentationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllInteractionPresentationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allInteractionPresentations",
		Method:             "GET",
		PathPattern:        "/workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllInteractionPresentationsReader{formats: a.formats},
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
	success, ok := result.(*AllInteractionPresentationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for allInteractionPresentations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteWorkflowUserInteractionPresentationExecution deletes workflow execution

Cancels a workflow execution user interaction presentation instance. This API call cancels only the workflow execution user interaction presentation. To cancel the workflow execution use APIs under <b>/api/workflows/{workflowId}/instances</b>
*/
func (a *Client) DeleteWorkflowUserInteractionPresentationExecution(params *DeleteWorkflowUserInteractionPresentationExecutionParams, opts ...ClientOption) (*DeleteWorkflowUserInteractionPresentationExecutionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkflowUserInteractionPresentationExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteWorkflowUserInteractionPresentationExecution",
		Method:             "DELETE",
		PathPattern:        "/workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWorkflowUserInteractionPresentationExecutionReader{formats: a.formats},
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
	success, ok := result.(*DeleteWorkflowUserInteractionPresentationExecutionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteWorkflowUserInteractionPresentationExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPresentationForInteraction gets presentation for interaction

Retrieves the definition of a user interaction presentation. To retrieve the definition of a user interaction presentation, make an HTTP GET request at the URL of the presentation. Pass the workflow and execution IDs as a path variable. To retrieve the user interaction presentation definition localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow of the user interaction, otherwise it defaults to the standard user interaction presentation definition.
*/
func (a *Client) GetPresentationForInteraction(params *GetPresentationForInteractionParams, opts ...ClientOption) (*GetPresentationForInteractionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPresentationForInteractionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPresentationForInteraction",
		Method:             "GET",
		PathPattern:        "/workflows/{workflowId}/executions/{executionId}/interaction/presentation",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPresentationForInteractionReader{formats: a.formats},
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
	success, ok := result.(*GetPresentationForInteractionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPresentationForInteraction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LoadExecutionForWorkflowUserInteractionPresentation loads execution

Retrieves a specific workflow execution user interaction presentation instance. To retrieve the user interaction presentation localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow of the user interaction, otherwise it defaults to the standard user interaction presentation.
*/
func (a *Client) LoadExecutionForWorkflowUserInteractionPresentation(params *LoadExecutionForWorkflowUserInteractionPresentationParams, opts ...ClientOption) (*LoadExecutionForWorkflowUserInteractionPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadExecutionForWorkflowUserInteractionPresentationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "loadExecutionForWorkflowUserInteractionPresentation",
		Method:             "GET",
		PathPattern:        "/workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LoadExecutionForWorkflowUserInteractionPresentationReader{formats: a.formats},
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
	success, ok := result.(*LoadExecutionForWorkflowUserInteractionPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for loadExecutionForWorkflowUserInteractionPresentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartInteractionPresentation starts interaction presentation

Creates a new instance of a workflow user interaction presentation by using the passed parameters. To create a new instance of a workflow user interaction presentation, make an HTTP GET request at the URL that contains the instances of that workflow user interaction presentation. Presentation fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid. In order the returned workflow user interaction presentation to be localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow, otherwise it defaults to the standard workflow user interaction presentation.
*/
func (a *Client) StartInteractionPresentation(params *StartInteractionPresentationParams, opts ...ClientOption) (*StartInteractionPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartInteractionPresentationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startInteractionPresentation",
		Method:             "POST",
		PathPattern:        "/workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartInteractionPresentationReader{formats: a.formats},
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
	success, ok := result.(*StartInteractionPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startInteractionPresentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateWorkflowUserInteractionPresentation updates presentation

Update a specific workflow execution user interaction presentation instance. Presentation fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid.
*/
func (a *Client) UpdateWorkflowUserInteractionPresentation(params *UpdateWorkflowUserInteractionPresentationParams, opts ...ClientOption) (*UpdateWorkflowUserInteractionPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkflowUserInteractionPresentationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateWorkflowUserInteractionPresentation",
		Method:             "POST",
		PathPattern:        "/workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkflowUserInteractionPresentationReader{formats: a.formats},
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
	success, ok := result.(*UpdateWorkflowUserInteractionPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateWorkflowUserInteractionPresentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateWorkflowUserInteractionPresentationPUT updates presentation

Update a specific workflow execution user interaction presentation instance. Presentation fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid.
*/
func (a *Client) UpdateWorkflowUserInteractionPresentationPUT(params *UpdateWorkflowUserInteractionPresentationPUTParams, opts ...ClientOption) (*UpdateWorkflowUserInteractionPresentationPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkflowUserInteractionPresentationPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateWorkflowUserInteractionPresentationPUT",
		Method:             "PUT",
		PathPattern:        "/workflows/{workflowId}/executions/{executionId}/interaction/presentation/instances/{presentationExecutionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkflowUserInteractionPresentationPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdateWorkflowUserInteractionPresentationPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateWorkflowUserInteractionPresentationPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
