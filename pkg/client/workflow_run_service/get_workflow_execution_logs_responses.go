// Code generated by go-swagger; DO NOT EDIT.

package workflow_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hmalphettes/vro-sdk-go/pkg/models"
)

// GetWorkflowExecutionLogsReader is a Reader for the GetWorkflowExecutionLogs structure.
type GetWorkflowExecutionLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowExecutionLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowExecutionLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkflowExecutionLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowExecutionLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/{workflowId}/executions/{executionId}/logs] getWorkflowExecutionLogs", response, response.Code())
	}
}

// NewGetWorkflowExecutionLogsOK creates a GetWorkflowExecutionLogsOK with default headers values
func NewGetWorkflowExecutionLogsOK() *GetWorkflowExecutionLogsOK {
	return &GetWorkflowExecutionLogsOK{}
}

/*
GetWorkflowExecutionLogsOK describes a response with status code 200, with default header values.

The request is successful
*/
type GetWorkflowExecutionLogsOK struct {
	Payload *models.Logs
}

// IsSuccess returns true when this get workflow execution logs o k response has a 2xx status code
func (o *GetWorkflowExecutionLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow execution logs o k response has a 3xx status code
func (o *GetWorkflowExecutionLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution logs o k response has a 4xx status code
func (o *GetWorkflowExecutionLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow execution logs o k response has a 5xx status code
func (o *GetWorkflowExecutionLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution logs o k response a status code equal to that given
func (o *GetWorkflowExecutionLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workflow execution logs o k response
func (o *GetWorkflowExecutionLogsOK) Code() int {
	return 200
}

func (o *GetWorkflowExecutionLogsOK) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/logs][%d] getWorkflowExecutionLogsOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowExecutionLogsOK) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/logs][%d] getWorkflowExecutionLogsOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowExecutionLogsOK) GetPayload() *models.Logs {
	return o.Payload
}

func (o *GetWorkflowExecutionLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Logs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowExecutionLogsUnauthorized creates a GetWorkflowExecutionLogsUnauthorized with default headers values
func NewGetWorkflowExecutionLogsUnauthorized() *GetWorkflowExecutionLogsUnauthorized {
	return &GetWorkflowExecutionLogsUnauthorized{}
}

/*
GetWorkflowExecutionLogsUnauthorized describes a response with status code 401, with default header values.

The user is not authorized
*/
type GetWorkflowExecutionLogsUnauthorized struct {
}

// IsSuccess returns true when this get workflow execution logs unauthorized response has a 2xx status code
func (o *GetWorkflowExecutionLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow execution logs unauthorized response has a 3xx status code
func (o *GetWorkflowExecutionLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution logs unauthorized response has a 4xx status code
func (o *GetWorkflowExecutionLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow execution logs unauthorized response has a 5xx status code
func (o *GetWorkflowExecutionLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution logs unauthorized response a status code equal to that given
func (o *GetWorkflowExecutionLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get workflow execution logs unauthorized response
func (o *GetWorkflowExecutionLogsUnauthorized) Code() int {
	return 401
}

func (o *GetWorkflowExecutionLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/logs][%d] getWorkflowExecutionLogsUnauthorized ", 401)
}

func (o *GetWorkflowExecutionLogsUnauthorized) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/logs][%d] getWorkflowExecutionLogsUnauthorized ", 401)
}

func (o *GetWorkflowExecutionLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowExecutionLogsNotFound creates a GetWorkflowExecutionLogsNotFound with default headers values
func NewGetWorkflowExecutionLogsNotFound() *GetWorkflowExecutionLogsNotFound {
	return &GetWorkflowExecutionLogsNotFound{}
}

/*
GetWorkflowExecutionLogsNotFound describes a response with status code 404, with default header values.

Can not find a workflow with the specified ID or the user does not have 'read' access rights for that workflow
*/
type GetWorkflowExecutionLogsNotFound struct {
}

// IsSuccess returns true when this get workflow execution logs not found response has a 2xx status code
func (o *GetWorkflowExecutionLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow execution logs not found response has a 3xx status code
func (o *GetWorkflowExecutionLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow execution logs not found response has a 4xx status code
func (o *GetWorkflowExecutionLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow execution logs not found response has a 5xx status code
func (o *GetWorkflowExecutionLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow execution logs not found response a status code equal to that given
func (o *GetWorkflowExecutionLogsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workflow execution logs not found response
func (o *GetWorkflowExecutionLogsNotFound) Code() int {
	return 404
}

func (o *GetWorkflowExecutionLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/logs][%d] getWorkflowExecutionLogsNotFound ", 404)
}

func (o *GetWorkflowExecutionLogsNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/{workflowId}/executions/{executionId}/logs][%d] getWorkflowExecutionLogsNotFound ", 404)
}

func (o *GetWorkflowExecutionLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
