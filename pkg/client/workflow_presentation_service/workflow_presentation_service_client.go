// Code generated by go-swagger; DO NOT EDIT.

package workflow_presentation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflow presentation service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow presentation service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllPresentation(params *AllPresentationParams, opts ...ClientOption) (*AllPresentationOK, error)

	DeleteWorkflowPresentationExecution(params *DeleteWorkflowPresentationExecutionParams, opts ...ClientOption) (*DeleteWorkflowPresentationExecutionNoContent, error)

	GetPresentationFor(params *GetPresentationForParams, opts ...ClientOption) (*GetPresentationForOK, error)

	LoadExecutionForWorkflowPresentation(params *LoadExecutionForWorkflowPresentationParams, opts ...ClientOption) (*LoadExecutionForWorkflowPresentationOK, error)

	StartPresentation(params *StartPresentationParams, opts ...ClientOption) (*StartPresentationOK, error)

	UpdateWorkflowPresentation(params *UpdateWorkflowPresentationParams, opts ...ClientOption) (*UpdateWorkflowPresentationOK, error)

	UpdateWorkflowPresentationNoResponse(params *UpdateWorkflowPresentationNoResponseParams, opts ...ClientOption) (*UpdateWorkflowPresentationNoResponseOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AllPresentation gets all presentations

Retrieves a list of the presentations for a workflow that you specify. To retrieve the list of workflow presentations, make an HTTP GET request at the workflow presentations list URL. The returned list contains all of the currently running workflow presentation instances, and all completed instances based on the data from the workflow executions. If the user has admin rights, all presentation instances for all users are returned.
*/
func (a *Client) AllPresentation(params *AllPresentationParams, opts ...ClientOption) (*AllPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllPresentationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allPresentation",
		Method:             "GET",
		PathPattern:        "/workflows/{workflowId}/presentation/instances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllPresentationReader{formats: a.formats},
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
	success, ok := result.(*AllPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for allPresentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteWorkflowPresentationExecution deletes workflow execution

Cancels the execution of a workflow presentation instance. This API call cancels only the workflow presentation execution. To cancel the workflow execution, use APIs under <b>/api/workflows/{workflowId}/instances</b>.
*/
func (a *Client) DeleteWorkflowPresentationExecution(params *DeleteWorkflowPresentationExecutionParams, opts ...ClientOption) (*DeleteWorkflowPresentationExecutionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkflowPresentationExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteWorkflowPresentationExecution",
		Method:             "DELETE",
		PathPattern:        "/workflows/{workflowId}/presentation/instances/{executionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWorkflowPresentationExecutionReader{formats: a.formats},
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
	success, ok := result.(*DeleteWorkflowPresentationExecutionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteWorkflowPresentationExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPresentationFor gets presentation

Retrieves the definition of a workflow presentation. To retrieve the workflow presentation definition localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow, otherwise it defaults to the standard workflow presentation definition.
*/
func (a *Client) GetPresentationFor(params *GetPresentationForParams, opts ...ClientOption) (*GetPresentationForOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPresentationForParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPresentationFor",
		Method:             "GET",
		PathPattern:        "/workflows/{workflowId}/presentation",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPresentationForReader{formats: a.formats},
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
	success, ok := result.(*GetPresentationForOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPresentationFor: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LoadExecutionForWorkflowPresentation loads execution

Retrieves a specific workflow presentation instance. Presentation instances are removed after the workflow starts. If the presentation instance under requested <b>executionId</b> does not exists, a new presentation instance is created by using the parameters from the workflow execution with the same ID. To retrieve the workflow presentation localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow, otherwise it defaults to the standard workflow presentation.
*/
func (a *Client) LoadExecutionForWorkflowPresentation(params *LoadExecutionForWorkflowPresentationParams, opts ...ClientOption) (*LoadExecutionForWorkflowPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadExecutionForWorkflowPresentationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "loadExecutionForWorkflowPresentation",
		Method:             "GET",
		PathPattern:        "/workflows/{workflowId}/presentation/instances/{executionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LoadExecutionForWorkflowPresentationReader{formats: a.formats},
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
	success, ok := result.(*LoadExecutionForWorkflowPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for loadExecutionForWorkflowPresentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartPresentation starts presentation

	Creates a new instance of the presentation of a workflow, by using the passed parameters. To create a new instance of a workflow presentation, make an HTTP GET request at the URL that contains the instances of the workflow presentation. Presentation's fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid. In order the returned workflow presentation to be localized, add Accept-Language header, with the appropriate locale. In advance, localization resource should be present for the workflow, otherwise it defaults to the standard workflow presentation.
*/
func (a *Client) StartPresentation(params *StartPresentationParams, opts ...ClientOption) (*StartPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPresentationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startPresentation",
		Method:             "POST",
		PathPattern:        "/workflows/{workflowId}/presentation/instances",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartPresentationReader{formats: a.formats},
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
	success, ok := result.(*StartPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startPresentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateWorkflowPresentation updates presentation

Update a specific workflow presentation instance. Presentation fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid. If the parameter's 'updated' flag is set to true, the dependent field values are recalculated.
*/
func (a *Client) UpdateWorkflowPresentation(params *UpdateWorkflowPresentationParams, opts ...ClientOption) (*UpdateWorkflowPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkflowPresentationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateWorkflowPresentation",
		Method:             "POST",
		PathPattern:        "/workflows/{workflowId}/presentation/instances/{executionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkflowPresentationReader{formats: a.formats},
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
	success, ok := result.(*UpdateWorkflowPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateWorkflowPresentation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateWorkflowPresentationNoResponse updates presentation

Update a specific workflow presentation instance. Presentation fields are populated with input parameter values and are validated. If there are any validation errors, they are collected and attached to each field. The presentation is marked as invalid. If the parameter's 'updated' flag is set to true, the dependent field values are recalculated.
*/
func (a *Client) UpdateWorkflowPresentationNoResponse(params *UpdateWorkflowPresentationNoResponseParams, opts ...ClientOption) (*UpdateWorkflowPresentationNoResponseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkflowPresentationNoResponseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateWorkflowPresentationNoResponse",
		Method:             "PUT",
		PathPattern:        "/workflows/{workflowId}/presentation/instances/{executionId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkflowPresentationNoResponseReader{formats: a.formats},
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
	success, ok := result.(*UpdateWorkflowPresentationNoResponseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateWorkflowPresentationNoResponse: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
